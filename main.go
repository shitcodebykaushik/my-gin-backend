package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Vercel expects an exported function called "Handler"
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()

	// Define the login endpoint
	router.POST("/login", loginHandler)

	// Serve the HTTP request
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

// Main function is not required for Vercel but can be included for local testing
func main() {
	// This won't be invoked by Vercel, but can be useful for local testing
}
