package books

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
var mu sync.Mutex

func InitData() {
	books = append(books, Book{
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

	json.NewEncoder(w).Encode(books)
}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	books = append(books, newBook)
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
	for _, book := range books {
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
	for i, book := range books {
		if book.ID == id {
			index = i
			break
		}

	}
	if index == -1 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	books = slices.Delete(books, index, index+1)
	w.WriteHeader(http.StatusNoContent)
}
