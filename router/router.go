package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	dataadapter "github.com/dilesh14/CeruleanSalween/db"
)

type CUserReq struct {
	Username    string `json:"uname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

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
		hContentType := r.Header.Get("Content-Type")
		if hContentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		var u CUserReq

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&u)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user := dataadapter.User{
			UserName:     u.Username,
			PhoneNumber:  u.PhoneNumber,
			EmailAddress: u.Email,
			Password:     u.Password,
		}
		uri := os.Getenv("mongoDbConnectionString")
		dataadapter.SetDbConnectionString(uri)
		dataadapter.CreateUser(user)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "user created"}`))
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
