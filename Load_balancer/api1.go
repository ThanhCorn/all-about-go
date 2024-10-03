package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Movies struct {
	Name       string `json:"name"`
	Release    int    `json:"release"`
	Collection string `json:"collection"`
}

// Hardcoded list of movies
var movies = []Movies{
	{Name: "Interstellar", Release: 2014, Collection: "1.23 Crores"},
	{Name: "Shutter Island", Release: 2010, Collection: "186.4 Crores"},
	{Name: "Tenet", Release: 2020, Collection: "20.4 Crores"},
	{Name: "The Prestige", Release: 2006, Collection: "185 Millions"},
	{Name: "Memento", Release: 2000, Collection: "9 Millions"},
	{Name: "Dunkirk", Release: 2017, Collection: "530 Millions"},
}

// GetMovies serves the static movie data
func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func main() {
	router := gin.Default()
	router.GET("/", GetMovies)
	router.Run("localhost:8081")
}
