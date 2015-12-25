package main

import (
	"log"
	"net/http"
	"time"
)

//var templates = template.Must(template.ParseFiles("checkin.html", "current.html"))

var timeout = 30 * time.Minute

func main() {
	OpenDB()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
