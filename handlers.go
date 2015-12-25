package main

import (
	"strings"
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
	_"github.com/go-sql-driver/mysql"
)

//Member contains a name that will be sent to the database
type Member struct {
	Name, CheckinTime string
}

//var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/checkin.html", "tmpl/current.html"))
var db *sql.DB

//IndexHandler handles calls to the index of the Server
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/ubuntu/gowrk/bin/static/index.html")
}

//CurrentHandler handles calls to the currently checked-in user list
func CurrentHandler(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM checkedin")
	defer rows.Close()

	var members []Member
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Name, &m.CheckinTime)
		members = append(members, m)
	}
	t, err := template.ParseFiles("/home/ubuntu/gowrk/bin/tmpl/current.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, members)
}

//Checkin adds team members to the database
func CheckinHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	now := time.Now().Format(time.Kitchen)
	res, err := db.Exec("INSERT INTO checkedin VALUES(?,?)", name, now)
	if err != nil {
		panic(err)
	}
	ra, _ := res.RowsAffected()
	fmt.Printf("Rows affected: %d", ra)
	time.AfterFunc(30*time.Minute, func() {
		db.Exec("DELETE FROM checkedin WHERE Name=?", name)
	})
	http.Redirect(w, r, "/", http.StatusFound)
}

//OpenDB opens up the database
func OpenDB() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Database name: ")
	database,_:=read.ReadString('\n')
	database=strings.Trim(database,"\n")
	fmt.Print("Username: ")
	user,_:=read.ReadString('\n')
	user=strings.Trim(user,"\n")
	fmt.Print("Password: ")
	password,_:=read.ReadString('\n')
	password=strings.Trim(password,"\n")
	var err error
	db, err = sql.Open("mysql", user+":"+password+"@/"+database+"?allowCleartextPasswords=true")
	if err != nil {
		panic(err)
	}

}
