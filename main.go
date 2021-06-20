package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the homepage!")
}

func handleRequest() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	handleRequest()
}
