package main

import (
	"booksMan/authorization"
	"booksMan/cors"
	"booksMan/handlers"
	"booksMan/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.DBConnSetup()

	router := gin.Default()
	cors.ConfigureCORS(router)
	router.Use(authorization.LoggerMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/gettotal", authorization.AuthMiddleware, handlers.FindTotal)
	router.POST("/allbook", authorization.AuthMiddleware, handlers.ListBooks)
	router.POST("/bookcreate", authorization.AuthMiddleware, handlers.AddBookHandler)
	router.POST("/book/search", authorization.AuthMiddleware, handlers.SearchBooksHandler)
	router.PATCH("/bookupdate", authorization.AuthMiddleware, handlers.UpdateBookHandler)
	router.DELETE("/book/:id", authorization.AuthMiddleware, handlers.DeleteBookHandler)
	// Protecting a route with the middleware

	router.POST("/borrowrecords", authorization.AuthMiddleware, handlers.ListBorrowRecordHandler)
	router.POST("/borrow", authorization.AuthMiddleware, handlers.IssueBookHandler)
	router.GET("/borrowedbooks", authorization.AuthMiddleware, handlers.SearchBorrowRecordHandler)
	router.POST("/borrow/return", authorization.AuthMiddleware, handlers.ReturnBookHandler)
	router.DELETE("/borrow/:id", authorization.AuthMiddleware, handlers.DeleteBorrowRecordHandler)
	router.Run()
}
