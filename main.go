package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/andrewtian/minepong"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	log.Fatalln(http.ListenAndServe(":8040", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	s := minepong.NewServer("desrtia", "pvp.desteria.com:25565")
	if err := s.Connect(); err != nil {
		fmt.Fprintln(w, "couldnt connect to server")
		return
	}

	pong, err := s.Ping()
	if err != nil {
		fmt.Println(w, "there was an error pinging sorry")
		return
	}

	fmt.Fprintf(w, "hello! %s has %d players of %d max", s.Name, pong.Players.Online, pong.Players.Max)
}
