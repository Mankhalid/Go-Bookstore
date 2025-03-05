package main

import (
	"database/sql"

	"log"
)

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("mysql", "root:081202@tcp(127.0.0.1:3306)/bookdb")
	if err != nil {
		log.Fatal("error connecting DB", err)
	}
	defer db.Close()
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}
