package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func newDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	checkErr(err)

	_, err = db.Exec("create table if not exists posts(title text, body text)")
	checkErr(err)

	return db
}

func main() {
	db := newDB()
	db.Close()
}
