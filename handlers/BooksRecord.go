package handlers

import (
	//	"booksMan/repository"

	"booksMan/models"
	"booksMan/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET all books
func ListBooks(c *gin.Context) {

	books, err := repository.GETreturnAllBooks()
	//fmt.Printf("books : %v", books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"books": books,
		"error": nil,
	})
}

// CREATE a book
func AddBookHandler(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := repository.POSTcreateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
			"book":  result,
		})
		return
	}
	c.JSON(200, gin.H{
		"error": nil,
		"book":  result,
	})
}

// Destroy a book
func DeleteBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		err = repository.DELETEdeleteBook(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

// Update a book
func UpdateBookHandler(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book) // Bind request body to User struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = repository.PATCHupdateBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"error": nil,
	})
}

// Search Books based on given  title
func SearchBooksHandler(c *gin.Context) {
	var books []models.Book
	type searchStr struct {
		searchString string
	}
	var s searchStr
	c.ShouldBindJSON(&s)
	length, err := repository.POSTsearchBook(books, s.searchString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"length": length,
		"books":  books,
		"error":  nil,
	})
}
