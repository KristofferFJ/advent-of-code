package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("code/snippetbox/ui/static/"))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
