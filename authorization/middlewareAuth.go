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
	//authHeader := c.Request.Header["Authorization"][0]
	print("Read request header\n")
	for key, value := range c.Request.Header {
		fmt.Println("Header:", key, "=", value)
	}
	print("\nHeader complete\n")
	//authHeader := "eedvdfbdbfdfbdfbdfbdfbdfbdbfbdfbdfbdb"
	if authHeader == "" {
		//fmt.Printf("Request :: \n%s\n", c.Request.Header["Authorization"][0])
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

	// upath := c.Request.URL.Path
	// upath = strings.Replace(upath, "/", "", 1)
	// body := struct {
	// 	Endpoint string `json:"endpoint"`
	// }{
	// 	Endpoint: upath,
	// }
	// json_body, err := json.Marshal(&body)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to create body"})
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	return
	// }

	req, err := http.NewRequest(http.MethodPost, authurl, nil)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Middleware : Could not create request"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Response status:", res.Status)
	fmt.Print("\n\n res ::: ", res)

	scanner := bufio.NewScanner(res.Body)
	var responcebody []string
	for i := 0; scanner.Scan(); i++ {
		responcebody = append(responcebody, (scanner.Text()))
		fmt.Printf("i %d  : responcebody[%d] :: %s\n", i, i, responcebody[i])
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("\nResponce body :\t")
	fmt.Println(responcebody)

	err = json.Unmarshal([]byte(responcebody[0]), &ResponceData)
	if err != nil {
		fmt.Print("Cannot unmarshal data : ", err)
	}
	fmt.Printf("\nCurrent User : %v\n\n", ResponceData)
	if ResponceData.Status == 200 {
		CurrentUser = ResponceData.User
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Un authorized",
		})
	}
}

// Create a decoder object from the response body

// Define a variable to store the decoded data
