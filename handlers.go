package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

//Member contains a name that will be sent to the database
type Member struct {
	Name        string
	CheckinTime time.Time
}

//var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/checkin.html", "tmpl/current.html"))
var db *sql.DB

type page struct {
	Name        string
	CheckInTime time.Time
}

var mainPage = page{Name: "Check In"}

//Index handles calls to the index of the Server
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

//Current handles calls to the currently checked-in user list
func Current(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/current.html")
	t.Execute(w, nil)
}

//Checkin adds team members to the database
func Checkin(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	now := time.Now().Format(time.Kitchen)
	res, err := db.Exec("INSERT INTO checkedin VALUES(?,?)", name, now)
	if err != nil {
		panic(err)
	}
	ra, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d", ra)
	//REMEMBER TO CHANGE THIS BACK TO 30 MINUTES!!!
	time.AfterFunc(30*time.Minute, func() {
		db.Exec("DELETE FROM checkedin WHERE Name=?", name)
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

//deleteEntry deletes a team member from the database.
//Only called after the 30 minute timer expires
func deleteEntry(name string) {

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
	db.Ping()
}

//CloseDB closes the connection to the sql server
func CloseDB() {
	db.Close()
}
