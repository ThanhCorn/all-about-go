package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Class   string `json:"Class"`
	Section string `json:"Section"`
}

// Hardcoded list of students
var students = []Student{
	{ID: "1ds23is023", Name: "Arpit_Srivastava", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is021", Name: "Anurag_Jain", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is022", Name: "Ananaya_Gowda", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is020", Name: "Khushi_Agrawal", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is024", Name: "Harsh_Gupta", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is025", Name: "Aaskar_Rana", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is026", Name: "Simran_Sharma", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is027", Name: "Nikhil_Kumar", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is028", Name: "Priya_Verma", Age: 20, Class: "3", Section: "A"},
	{ID: "1ds23is029", Name: "Rohan_Mehra", Age: 20, Class: "3", Section: "A"},
}

// GetAllStudentsNames serves the static student data
func GetAllStudentsNames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func main() {
	router := gin.Default()
	router.GET("/", GetAllStudentsNames)
	router.Run("localhost:8083")
}
