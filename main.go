package main

import (
	"time"
	"net/http"
	"encoding/json"

)

type Book struct {
	ID				string		`json:"id"`
	Title			string		`json:"title"`
	Author			string		`json:"author"`
	PublishedDate	time.Time	`json:"published_date"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	json.NewEncoder(w).Encode(books)
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var newBook Book
	err :=json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	books = append(books,newBook)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)



}