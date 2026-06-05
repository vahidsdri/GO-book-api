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