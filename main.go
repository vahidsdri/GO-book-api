package main

import (
	"net/http"
	"go-book-api/books"
)


func main() {
	books.InitData()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", books.GetBooks)
	mux.HandleFunc("POST /books", books.CreateBooks)
	mux.HandleFunc("GET /books/{id}", books.GetBookById)
	mux.HandleFunc("DELETE /books/{id}", books.DeleteBook)

	println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
