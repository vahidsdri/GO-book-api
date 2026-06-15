package models

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Book struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

var DB *sql.DB

func InitDB(){
	var err error

	DB, err = sql.Open("sqlite", "books.db")
	if err != nil{
		log.Fatal("Failed to connect to database.", err)
	}
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS books (
		id TEXT PRIMARY KEY,
		title TEXT,
		author TEXT,
		published_date DATETIME
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Faild to create table.", err)
	}
	log.Println("Database initialized successfully!")
}


