package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates = make(map[string]*template.Template)

func init() {
	templates["list.html"] = template.Must(template.ParseFiles("templates/list.tmpl", "templates/layout.tmpl"))
	templates["lisdt.html"] = template.Must(template.ParseFiles("templates/list.tmpl", "templates/layout.tmpl"))
	templates["about.html"] = template.Must(template.ParseFiles("templates/about.tmpl", "templates/layout.tmpl"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ListHandler)
	r.HandleFunc("/test", TestHandler)
	r.HandleFunc("/about", AboutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
