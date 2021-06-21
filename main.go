package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dataadapter "github.com/dilesh14/CeruleanSalween/db"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage! \n This website will change over the course of time as i get familiar with web programming with go. \n Oh Yeah! the backend of this website is written using go.")
}

func getDb(w http.ResponseWriter, r *http.Request) {
	uri := os.Getenv("mongoDbConnectionString")
	if uri == "" {
		log.Fatalf("db uri found not")
	}

	databases := dataadapter.LoadAllDatabases(uri)
	fmt.Fprintf(w, "databases: %v\n", databases)
}

func handleRequest() {
	envVarError := godotenv.Load()
	if envVarError != nil {
		log.Fatal("Error loading .env file", envVarError)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/db", getDb)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequest()
}
