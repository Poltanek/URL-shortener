/*
ToDo:
Go HTTP Server
Concurrency-safe data structures
JSON handling
Routing and request handling
Basic URL shortening logic
*/


package main

// fml - for printing to the console
// log - is for logging errors
// net/http - is for building the web server

import (
	"fml"
	"log"
	"net/http"
)



func main() {
	store := NewURLStore()

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		ShortenURLHandler(w, r, store)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.request) {
		RedirectHandler(w, r, store)
	})

	port := ":8000"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}