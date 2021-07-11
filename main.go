package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dilesh14/CeruleanSalween/router"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.HandleRoutes()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
