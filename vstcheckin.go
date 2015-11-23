import (
  "time"
  "sql"
  "regexp"
  "errors"
  "net/http"
  "html/template"
  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/native"

)
var templates = template.Must(template.ParseFiles("checkin.html","current.html"))
var paths=regexp.MustCompile("^/(checkin|current))$")

var timeout = 30 * time.Minute
func main(){
  db:=mysql.New("tcp","",172.0.0.1:3306,"moatman","nope")
  err:=db.Connect()
  if err != nil{
    panic(err)
  }


}
