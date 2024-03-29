package main

import (
	"errors"
	"github.com/andrewtian/minepong"
	"github.com/dustin/go-humanize"
	"net"
	"strconv"
	"time"
)

type ServerStatus int

const (
	Active  ServerStatus = iota + 1 // 1
	Unknown                         // 2
)

// the schedule of what to ping each server
// Active: the server responded to our last request
// Unknown: the server did not
var pSchedule = map[ServerStatus]time.Duration{
	Active:  10 * time.Second,
	Unknown: 60 * time.Second,
}

var lm = NewListManager(pSchedule)

type Server struct {
	Name        string
	Host        string
	Conn        net.Conn
	PongHistory []*Pong
	Status      ServerStatus
}

type Pong struct {
	minepong.Pong
	latency time.Duration
	time    time.Time
}

func NewServer(name string, host string) *Server {
	return &Server{
		Name:        name,
		Host:        host,
		PongHistory: []*Pong{},
		Status:      Active,
	}
}

func (s *Server) Connect() error {
	var err error
	s.Conn, err = net.Dial("tcp", s.Host)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Disconnect() {
	s.Conn.Close()
}

func (s *Server) LastPing() *Pong {
	if len(s.PongHistory) > 0 {
		return s.PongHistory[0]
	}

	return nil
}

func (s *Server) Latency() string {
	if s.LastPing() == nil {
		return "n/a"
	}

	return s.LastPing().latency.String()
}

func (s *Server) Players() string {
	if s.LastPing() == nil {
		return "n/a"
	}

	return strconv.Itoa(s.LastPing().Players.Online) + "/" + strconv.Itoa(s.LastPing().Players.Max)
}

func (s *Server) LastPingTime() string {
	if s.LastPing() == nil {
		return "n/a"
	}

	return humanize.Time(s.LastPing().time)
}

func (s *Server) Ping() error {
	ts := time.Now()

	if err := s.Connect(); err != nil {
		return err
	}
	defer s.Disconnect()

	pong, err := minepong.Ping(s.Conn, s.Host)
	if err != nil {
		return errors.New("jjhkjhg")
	}

	lat := time.Now().Sub(ts)

	opong := &Pong{
		Pong:    *pong,
		latency: lat,
		time:    ts,
	}

	s.PongHistory = append([]*Pong{opong}, s.PongHistory...)
	return nil
}

type ListManager struct {
	servers      []*Server
	pingSchedule map[ServerStatus]time.Duration
}

func NewListManager(ps map[ServerStatus]time.Duration) *ListManager {
	return &ListManager{
		servers:      []*Server{},
		pingSchedule: ps,
	}
}

func (m *ListManager) Start() {
	go m.PingInterval(Unknown)
	go m.PingInterval(Active)
}

func (m *ListManager) PingInterval(typ ServerStatus) {
	for {
		m.PingType(typ)
		time.Sleep(pSchedule[typ])
	}
}

func (m *ListManager) PingAll() {
	m.PingType(Unknown)
	m.PingType(Active)
}

func (m *ListManager) PingType(typ ServerStatus) {
	for _, s := range m.servers {
		if s.Status != typ {
			continue
		}

		if err := s.Ping(); err != nil {
			s.Status = Unknown
			continue
		}

		// responded, set back to active status
		if s.Status == Unknown {
			s.Status = Active
		}
	}
}

func (m *ListManager) AddServer(s *Server) {
	m.servers = append(m.servers, s)
}
