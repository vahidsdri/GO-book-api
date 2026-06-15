package main

import (
	"net/http"
	"go-book-api/handlers"
)


func main() {
	handlers.InitData()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", handlers.GetBooks)
	mux.HandleFunc("POST /books", handlers.CreateBooks)
	mux.HandleFunc("GET /books/{id}", handlers.GetBookById)
	mux.HandleFunc("DELETE /books/{id}", handlers.DeleteBook)

	println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
