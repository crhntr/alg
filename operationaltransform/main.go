package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(w, r, hub)
	})

	http.HandleFunc("/", serveIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
