package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/checkin.html", "tmpl/current.html"))
var db *sql.DB

type page struct {
	Name string
}

var mainPage = page{Name: "Check In"}

//Index handles calls to the index of the Server
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

//Current handles calls to the currently checked-in user list
func Current(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl\\current.html")
	t.Execute(w, nil)
}

//Checkin adds team members to the database
func Checkin(w http.ResponseWriter, r *http.Request) {
	name := "Moatman"
	db.Exec("INSERT|vst|name=?,time=NOW", name)
}

//OpenDB the database
func OpenDB() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Database name: ")
	database, _ := read.ReadString('\n')
	fmt.Print("Username: ")
	user, _ := read.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := read.ReadString('\n')
	var err error
	db, _ = sql.Open("mymysql", database+"/"+user+"/"+password)
	if err != nil {
		panic(err)
	}
}
