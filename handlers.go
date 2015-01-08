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
	templates["list.html"].ExecuteTemplate(w, "base", map[string]interface{}{})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	templates["about.html"].ExecuteTemplate(w, "base", map[string]interface{}{})
}
