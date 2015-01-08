package main

import (
	"fmt"
	"github.com/andrewtian/minepong"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	s := NewServer("desrtia", "pvp.desteria.com:25565")
	if err := s.Connect(); err != nil {
		fmt.Fprintln(w, "couldnt connect to server")
		return
	}

	pong, err := minepong.Ping(s.Conn, s.Host)
	if err != nil {
		fmt.Println(w, "there was an error pinging sorry")
		return
	}

	fmt.Fprintf(w, "hello! %s has %d players of %d max", s.Name, pong.Players.Online, pong.Players.Max)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"servers": lm.servers,
	}

	templates["list.html"].ExecuteTemplate(w, "base", data)
}

func NewTrackHandler(w http.ResponseWriter, r *http.Request) {
	templates["track.html"].ExecuteTemplate(w, "base", map[string]interface{}{})
}

func CreateTrackHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	if err := r.ParseForm(); err != nil {
		templates["track.html"].ExecuteTemplate(w, "base", data)
	}

	name := r.Form.Get("name")
	if !(len(name) <= 64 && len(name) > 0) {
		data["error"] = "name must be between 0 and 65 characters"
		templates["track.html"].ExecuteTemplate(w, "base", data)
		return
	}

	host := r.Form.Get("hostname") + ":" + r.Form.Get("port")
	svr := NewServer(name, host)
	if err := svr.Ping(); err != nil {
		data["error"] = "could not verify server. check hostname and port"
		templates["track.html"].ExecuteTemplate(w, "base", data)
		return
	}

	lm.AddServer(svr)

	templates["track.html"].ExecuteTemplate(w, "base", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	templates["about.html"].ExecuteTemplate(w, "base", map[string]interface{}{})
}
