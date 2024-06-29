package authorization

import (
	"booksMan/models"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var CurrentUser models.User

type response struct {
	User   models.User `json:"user"`
	Status int         `json:"status"`
}

var ResponceData response

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before the handler")
		c.Next()
		fmt.Println("After the handler")
	}
}

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var authurl string = "http://localhost:3001/auth_controller"
	// Create a new HTTP client
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, authurl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Middleware : Could not create request"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error recided from client.Do \n err :\t%s\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)

	}
	defer res.Body.Close()
	//fmt.Print("\n\n res ::: ", res)

	scanner := bufio.NewScanner(res.Body)
	var responcebody []string
	for i := 0; scanner.Scan(); i++ {
		responcebody = append(responcebody, (scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if res.Status == "200 OK" {
		err = json.Unmarshal([]byte(responcebody[0]), &ResponceData)
		if err != nil {
			fmt.Print("Cannot unmarshal data : ", err)
		}
		fmt.Printf("\nCurrent User : %v\n\n", ResponceData)
		if ResponceData.Status == 200 {
			CurrentUser = ResponceData.User
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Un authorized",
			})
			c.AbortWithStatus(http.StatusForbidden)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": responcebody[0],
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
