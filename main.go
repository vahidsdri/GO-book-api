package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"sync"
	"time"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

var books []Book
var Mu sync.Mutex
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Mu.Lock()
	defer Mu.Unlock()

	json.NewEncoder(w).Encode(books)
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Mu.Lock()
	books = append(books, newBook)
	Mu.Unlock()
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newBook)

}
func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Extracting ID from the URL path
	id := r.PathValue("id")
	//Looping through slices of books
	Mu.Lock()
	defer Mu.Unlock()
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	//If the loop finishes and we haven't returned, the book doesn't exist
	http.Error(w, "Book not found.", http.StatusNotFound)

}

func deleteBook(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")

	if id == ""{
		http.Error(w,"Missing book ID.",http.StatusBadRequest)
		return
	}

	index := -1
	Mu.Lock()
	defer Mu.Unlock()
	for i, book := range books{
		if book.ID == id{
			index = i
			break
		}
		
	}
	if index == -1{
			http.Error(w,"Book not found", http.StatusNotFound)
			return
		}
	books = slices.Delete(books, index, index+1)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	books = append(books, Book{
		ID:            "1",
		Title:         "The Go Programming Language",
		Author:        "Meee",
		PublishedDate: time.Now(),
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", getBooks)
	mux.HandleFunc("POST /books", createBooks)
	mux.HandleFunc("GET /books/{id}", getBookById)
	mux.HandleFunc("DELETE /books/{id}", deleteBook)

	println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
