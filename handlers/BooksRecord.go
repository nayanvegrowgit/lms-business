package handlers

import (
	//	"booksMan/repository"

	"booksMan/authorization"
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
	if authorization.CurrentUser.Role_ID != 3 {

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
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un Authorized",
		})
		return
	}

}

// Destroy a book
func DeleteBookHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID != 3 {

		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			err = repository.DELETEdeleteBook(uint(id))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			c.JSON(http.StatusNoContent, nil)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un Authorized",
		})
		return
	}
}

// Update a book
func UpdateBookHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID != 3 {
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
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un Authorized",
		})
		return
	}
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
