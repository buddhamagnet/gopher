package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Page struct{}

func main() {
	r := mux.NewRouter()
	staticHandler := http.FileServer(http.Dir("./public/"))
	http.Handle("/public/", http.StripPrefix("/public/", staticHandler))
	r.HandleFunc("/", Root)
	http.Handle("/", r)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Root is the hypermedia root endpoint of the service (GET).
func Root(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, &Page{})
}
