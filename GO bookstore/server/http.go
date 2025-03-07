package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func addBook(w http.ResponseWriter, r *http.Request) {

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /books")
		addBook(w, r)
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request on /test")
	})

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE isbn=?)", book.ISBN).Scan(&exists)
	if err != nil {
		http.Error(w, "error checking ISBN", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "this is a duplicate isbn", http.StatusConflict)
		return
	}
	_, err = db.Exec("INSERT INTO books (title,author,isbn) VALUES (?,?,?)", book.Title, book.Author, book.ISBN)
	if err != nil {
		http.Error(w, "error inserting book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
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

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Title, &book.Author, &book.ISBN)
		if err != nil {
			http.Error(w, "Error scanning book", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
