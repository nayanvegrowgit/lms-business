package handlers

import (
	"booksMan/authorization"
	"booksMan/models"
	"booksMan/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListBorrowRecordHandler(c *gin.Context) {
	if authorization.CurrentUser.Role_ID != 3 {
		brs, err := repository.AllBorrowRecord()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"borrowing_record": brs,
			"error":            nil,
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
	//var result []models.BorrowingRecordResult
	if authorization.CurrentUser.Role_ID == 3 {
		fmt.Printf("user id in SearchBorrowRecordHandler :: %d\n", authorization.CurrentUser.ID)
		var err error
		result, err := repository.SearchBorrowRecord(authorization.CurrentUser.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"borrowing_record": nil,
				"error":            err,
			})
			return
		}
		c.JSON(200, gin.H{
			"borrowing_record": result,
			"error":            nil,
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
		type requestBody struct {
			Id uint `json:"id"`
		}
		var s requestBody
		err := c.ShouldBindJSON(&s)
		fmt.Printf("\nrequest body : %v\n", s)
		fmt.Printf("bind json err %v", err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = repository.UpdateBorrowRecord(s.Id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
