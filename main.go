package main

import (
	"fmt"
	"log"
	"os"

	"github.com/imsumedhaa/URL-Shortner/api"
	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No env file found")
	}
}

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || username == "" || password == "" || dbname == "" {
		log.Fatal("Missing one or more required environment variables")
	}

	httpConfig, err := api.NewHttp(host, port, username, password, dbname)

	if err != nil {
		fmt.Printf("Error creating the http connection: %v\n", err)
		os.Exit(1)
	}

	if err := httpConfig.Run(); err != nil {
		fmt.Printf("Error run http server: %v\n", err)
		os.Exit(1)
	}

}
