package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Quality int    `json:"quality"`
}

// Hardcoded list of books
var books = []book{
	{Title: "Hamlet", Author: "William Shakespeare", Quality: 2},
	{Title: "Pride and Prejudice", Author: "Jane Austen", Quality: 3},
	{Title: "1984", Author: "George Orwell", Quality: 4},
	{ID: "7", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quality: 3},
	{Title: "Moby-Dick", Author: "Herman Melville", Quality: 2},
	{Title: "To Kill a Mockingbird", Author: "Harper Lee", Quality: 5},
	{ID: "10", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quality: 4},
	{Title: "War and Peace", Author: "Leo Tolstoy", Quality: 4},
	{Title: "The Odyssey", Author: "Homer", Quality: 5},
	{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Quality: 4},
}

// getBooks serves the static book data
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/", getBooks)
	router.Run("localhost:8082")
}
