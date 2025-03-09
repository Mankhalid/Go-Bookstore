package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func addBook(w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /test")
	})

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book Book
	//
	err := json.NewDecoder(r.Body).Decode(&book)
	//
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	//
	err = InsertBook(db, book, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	//
}

func ListBooks(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT title, author, isbn FROM books")
	if err != nil {
		http.Error(w, "Error fetching books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	//
	books, err := FetchBooks(rows, w)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
