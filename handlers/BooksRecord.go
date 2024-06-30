package handlers

import (
	//	"booksMan/repository"

	"booksMan/authorization"
	"booksMan/models"
	"booksMan/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET total number of books ::
func FindTotal(c *gin.Context) {
	result, err := repository.GetTotalAvailableBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"sum":   result[0],
			"error": nil,
		})
	}
}

// GET all books
func ListBooks(c *gin.Context) {
	type Constraint struct {
		Offset uint   `json:"offset"`
		Limit  uint   `json:"limit"`
		Filter string `json:"filter`
	}
	var constraint Constraint

	err := c.ShouldBindJSON(&constraint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("constraint : %v\n", constraint)

	books, err := repository.GETreturnAllBooks(constraint.Offset, constraint.Limit, constraint.Filter)
	//fmt.Printf("books : %v", books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	fmt.Printf("Books : %v\n", books)
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
		fmt.Printf("book : %v\n", book)

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
