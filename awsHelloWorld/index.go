package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "awuser:lopolopo@tcp(mydbinstance.cwizkxtcxkss.us-east-2.rds.amazonaws.com:3306)"+
		"/test?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Println(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/test", testRoute)
	http.HandleFunc("/friends", getFriends)
	http.HandleFunc("/createNewFriend", createNewFriend)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := db.Ping()
	if err != nil {
		log.Println(err)
	}
	_, err = fmt.Fprint(w, "<h1> Fikusu... dnes dota?</h1>", err)
	if err != nil {
		log.Println(err)
	}
}

func getFriends(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT name FROM friends")
	if err != nil {
		log.Println(err)
	}
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
		}
		s += name + "\n"
	}
	_, err = fmt.Fprintln(w, s)
	if err != nil {
		log.Println(err)
	}
}

func createNewFriend(w http.ResponseWriter, _ *http.Request) {
	stmt, err := db.Prepare("INSERT into friends (id, name) value ('8', 'Nikols')")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}
	_, err = fmt.Fprintln(w, "ADDED NEW FRIEND")
	if err != nil {
		log.Println(err)
	}
}

func testRoute(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "<p>test route</p>")
	if err != nil {
		log.Println(err)
	}
}
