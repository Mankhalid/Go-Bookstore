package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initDB()
	defer db.Close()

	log.Println("registering routes")

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /books")
		addBook(w, r)
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /test")
	})

	log.Println("Server is running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
