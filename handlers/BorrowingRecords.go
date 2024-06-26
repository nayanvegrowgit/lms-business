package handlers

import (
	"booksMan/authorization"
	"booksMan/models"
	"booksMan/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBorrowRecordHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID != 3 {
		var brs []models.BorrowingRecord

		brs, err := repository.AllBorrowRecord()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"brs":   brs,
			"error": nil,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un Authorized Member Trying to access all borrow records.",
		})
		return
	}
}
func IssueBookHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID == 3 {
		var br models.BorrowingRecord
		err := c.ShouldBindJSON(&br)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := repository.CreateBorrowRecord(&br)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
				"br":    result,
			})
			return
		}
		c.JSON(200, gin.H{
			"error": nil,
			"br":    result,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un Authorized IssueBookHandler user is not member",
		})
		return
	}
}
func SearchBorrowRecordHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID == 3 {
		var brs []models.Book

		brs, err := repository.SearchBorrowRecord(authorization.CurrentUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"books": nil,
				"error": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"books": brs,
			"error": nil,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"books": nil,
			"error": "Un Authorized Login with member previledges",
		})
		return
	}
}

func ReturnBookHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID == 3 {
		var br models.BorrowingRecord
		type requestBody struct {
			id uint
		}
		var s requestBody
		err := c.ShouldBindJSON(&s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		br, err = repository.UpdateBorrowRecord(s.id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"br":    br,
			"error": nil,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"br":    nil,
			"error": "Un Authorized",
		})
		return
	}
}
func DeleteBorrowRecordHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID != 3 {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			err = repository.DeleteBorrowRecord(uint(id))
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
