package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

var books = []book{
	{ID: "1", Title: "Of Mice and Men", Author: "John Steinbeck", Genre: "fiction"},
	{ID: "2", Title: "The Pearl", Author: "John Steinbeck", Genre: "fiction"},
	{ID: "3", Title: "East of Eden", Author: "John Steinbeck", Genre: "fiction"},
	{ID: "4", Title: "The Grapes of Wrath", Author: "John Steinbeck", Genre: "fiction"},
	{ID: "5", Title: "Economics in One Lesson", Author: "Henry Hazlitt", Genre: "nonfiction"},
	{ID: "6", Title: "Mayflower", Author: "Nathaniel Philbrick", Genre: "nonfiction"},
	{ID: "7", Title: "The Problem of Pain", Author: "CS Lewis", Genre: "nonfiction"},
	{ID: "8", Title: "Surprised by Joy", Author: "CS Lewis", Genre: "nonfiction"},
	{ID: "9", Title: "Till We Have Faces", Author: "CS Lewis", Genre: "fiction"},
	{ID: "10", Title: "Oliver Twist", Author: "Charles Dickens", Genre: "fiction"},
	{ID: "11", Title: "Animal Farm", Author: "George Orwell", Genre: "fiction"},
	{ID: "12", Title: "Nineteen Eighty-Four", Author: "George Orwell", Genre: "fiction"},
}

func handler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Gabriel's Library API :)\nhttps://github.com/gabetrau/cyderes_skills_challenge")
}

func sortBooks(titles []string, books []book) []book {
	sort.Strings(titles)
	var sortedBooks = []book{}
	var size = len(books)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if strings.EqualFold(titles[i], books[j].Title) {
				sortedBooks = append(sortedBooks, books[j])
				continue
			}
		}
	}
	return sortedBooks
}

func getBooks(c *gin.Context) {
	var titles = []string{}
	for _, a := range books {
		titles = append(titles, a.Title)
	}
	books = sortBooks(titles, books)
	c.IndentedJSON(http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func formatAuthor(a string) string {
	return strings.ReplaceAll((strings.ReplaceAll(a, " ", "")), ".", "")
}

func getBookByAuthor(c *gin.Context) {
	author := c.Param("author")
	var foundBooks = []book{}
	var titles = []string{}
	for _, a := range books {
		if strings.EqualFold(formatAuthor(a.Author), formatAuthor(author)) {
			foundBooks = append(foundBooks, a)
			titles = append(titles, a.Title)
		}
	}
	if len(foundBooks) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books by that author found"})
		return
	}
	foundBooks = sortBooks(titles, foundBooks)
	c.IndentedJSON(http.StatusOK, foundBooks)
}

func main() {
	// Determine port for HTTP service.
	log.Print("starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	router := gin.Default()
	router.GET("/", handler)
	router.GET("/books", getBooks)
	router.GET("/books/:author", getBookByAuthor)
	router.POST("/books", postBooks)

	router.Run(":" + port)
}
