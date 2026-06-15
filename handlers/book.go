package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"sync"
	"time"
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
	"go-book-api/models"

)



var mu sync.Mutex

func InitData() {
	models.Books = append(models.Books, models.Book{
		ID:            "1",
		Title:         "The Go Programming Language",
		Author:        "Meee",
		PublishedDate: time.Now(),
	})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()

	json.NewEncoder(w).Encode(models.Books)
}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	models.Books = append(models.Books, newBook)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newBook)

}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Extracting ID from the URL path
	id := r.PathValue("id")
	//Looping through slices of books
	mu.Lock()
	defer mu.Unlock()
	for _, book := range models.Books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	//If the loop finishes and we haven't returned, the book doesn't exist
	http.Error(w, "Book not found.", http.StatusNotFound)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "Missing book ID.", http.StatusBadRequest)
		return
	}

	index := -1
	mu.Lock()
	defer mu.Unlock()
	for i, book := range models.Books {
		if book.ID == id {
			index = i
			break
		}

	}
	if index == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	models.Books = slices.Delete(models.Books, index, index+1)
	w.WriteHeader(http.StatusNoContent)
}
