package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func main() {
	serverURL := "http://localhost:9090/books"

	var book Book
	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&book.Title)

	fmt.Print("Enter Author Name: ")
	fmt.Scanln(&book.Author)

	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&book.ISBN)

	jsonData, err := json.Marshal(book)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Book successfully registered!")
	} else {
		fmt.Printf("Failed to register book. Status Code: %d\n", resp.StatusCode)
	}
}
