package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a simple handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	// Server runs on port 8080
	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
