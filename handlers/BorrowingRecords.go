package handlers

import (
	"booksMan/models"
	"booksMan/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBorrowRecordHandler(c *gin.Context) {
	var brs []models.BorrowingRecord

	length, err := repository.AllBorrowRecord(brs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"length": length,
		"brs":    brs,
		"error":  nil,
	})
}
func IssueBookHandler(c *gin.Context) {
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
}
func SearchBorrowRecordHandler(c *gin.Context) {
	var brs []models.BorrowingRecord
	type requestBody struct {
		userID uint
	}
	var s requestBody
	c.ShouldBindJSON(&s)
	length, err := repository.SearchBorrowRecord(brs, s.userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"length": length,
		"brs":    brs,
		"error":  nil,
	})
}

func ReturnBookHandler(c *gin.Context) {
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
}
func DeleteBorrowRecordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		err = repository.DeleteBorrowRecord(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
