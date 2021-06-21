package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage! \n This website will change over the course of time as i get familiar with web programming with go. \n Oh Yeah! the backend of this website is written using go.")
}

func handleRequest() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	envVarError := godotenv.Load()
	if envVarError != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequest()
}
