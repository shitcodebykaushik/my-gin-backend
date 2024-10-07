package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Login endpoint
	r.POST("/login", loginHandler)

	// Start the server
	r.Run() // listen and serve on the default port
}

// Handler for login
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
