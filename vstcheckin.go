package main

import (
	"bufio"
	"fmt"
	"http"
	"os"
	"sql"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/ziutek/mymysql/godrv"
)

var templates = template.Must(template.ParseFiles("checkin.html", "current.html"))

var timeout = 30 * time.Minute

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Database name: ")
	database, _ := read.ReadString('\n')
	fmt.Print("Username: ")
	user, _ := read.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := read.ReadString('\n')
	db, err := sql.Open("mymysql", database+"/"+user+"/"+password)
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.fatal(http.ListenAndServe(":8080", router))

}
