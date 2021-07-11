package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	dataadapter "github.com/dilesh14/CeruleanSalween/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello to the homepage! \n This website will change over the course of time as i get familiar with web programming with go. \n Oh Yeah! the backend of this website is written using go.")
}

func getDb(w http.ResponseWriter, r *http.Request) {
	uri := os.Getenv("mongoDbConnectionString")
	if uri == "" {
		log.Fatalf("db uri found not")
	}

	databases := dataadapter.LoadAllDatabases(uri)
	fmt.Fprintf(w, "databases: %v\n", databases)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		user := dataadapter.User{
			UserName:     "buklau",
			PhoneNumber:  "123131321",
			EmailAddress: "test.go.go",
			Password:     "this is yo ma",
		}
		uri := os.Getenv("mongoDbConnectionString")
		dataadapter.SetDbConnectionString(uri)
		dataadapter.CreateUser(user)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func HandleRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/db", getDb)
	http.HandleFunc("/users/create", createUser)
	fs := http.FileServer(http.Dir("content/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
