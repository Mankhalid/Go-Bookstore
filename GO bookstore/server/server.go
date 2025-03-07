package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var w http.ResponseWriter
var r *http.Request

func main() {
	initDB()
	defer db.Close()
	var choice int
	fmt.Println("press 1 for adding a book and 2 for listing all books")
	fmt.Scanln(&choice)
	if choice == 1 {
		addBook(w, r)
		log.Println("registering routes")
	} else {
		ListBooks(w, r)
		log.Println("listing books...")
	}
	log.Println("Server is running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
