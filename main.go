package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dilesh14/CeruleanSalween/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.HandleRoutes()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
