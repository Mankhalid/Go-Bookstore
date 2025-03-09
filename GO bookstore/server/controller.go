package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func InsertBook(db *sql.DB, book Book, w http.ResponseWriter) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE isbn=?)", book.ISBN).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking ISBN: %v", err)
	}
	if exists {
		http.Error(w, "this is a duplicate isbn", http.StatusConflict)
		return nil
	}

	_, err = db.Exec("INSERT INTO books (title, author, isbn) VALUES (?, ?, ?)", book.Title, book.Author, book.ISBN)
	if err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}
func FetchBooks(rows *sql.Rows, w http.ResponseWriter) ([]Book, error) {
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Title, &book.Author, &book.ISBN)
		if err != nil {
			http.Error(w, "Error scanning book", http.StatusInternalServerError)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
