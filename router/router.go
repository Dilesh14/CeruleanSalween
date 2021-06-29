package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dataadapter "github.com/dilesh14/CeruleanSalween/db"
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

func HandleRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/db", getDb)
}
