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
	Name, CheckinTime string
}

//var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/checkin.html", "tmpl/current.html"))
var db *sql.DB

//Index handles calls to the index of the Server
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

//Current handles calls to the currently checked-in user list
func Current(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM checkedin")
	defer rows.Close()

	var members []Member
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Name, &m.CheckinTime)
		members = append(members, m)
	}
	t, err := template.ParseFiles("./tmpl/current.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, members)
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

//OpenDB the database
func OpenDB() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Database name: ")
	database, _ := read.ReadString('\n')
	fmt.Print("Username: ")
	user, _ := read.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := read.ReadString('\n')
	db, _ = sql.Open("mymysql", database+"/"+user+"/"+password)
	var err error
	//db, _ = sql.Open("mymysql", "vst/newuser/")
	if err != nil {
		panic(err)
	}

}
