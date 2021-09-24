package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = ":8000"
)

type Person struct {
	Name string
	Id   string
}

func rendertemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "Samuel", Id: "12345"}
	parsedTemplate, _ := template.ParseFiles("templates/fitemplate.html")
	err := parsedTemplate.Execute(w, person)

	if err != nil {
		log.Println("Error executing template ", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rendertemplate).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.HandleFunc("/", rendertemplate)
	err := http.ListenAndServe(CONN_HOST+CONN_PORT, router)
	if err != nil {
		log.Fatal("Error starting up the server")
		return
	}
}
