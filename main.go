package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.POST("/login", loginHandler)

	// Serve HTTP requests
	router.ServeHTTP(w, r)
}

func loginHandler(c *gin.Context) {
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Replace with your predefined credentials
	if json.Username == "admin" && json.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"status": "logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}

func main() {
	// The main function is still required but won't serve as an entry point for Vercel.
	fmt.Println("This function is running on Vercel!")
}
