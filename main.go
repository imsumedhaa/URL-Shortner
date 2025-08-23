package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/imsumedhaa/URL-Shortner/api"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || username == "" || password == "" || dbname == "" {
		log.Fatal("Missing one or more required environment variables")
	}
	_, err := api.NewHttp(host, port, username, password, dbname)

	if err != nil {
		fmt.Printf("Error creating the http connection: %v\n", err)
		os.Exit(1)
	}


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
