package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("checkin.html", "current.html"))

//Index handles calls to the index of the Server
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.printf("tst")
}
func render(w http.ResponseWriter, tmpl string) {
	err := template.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
