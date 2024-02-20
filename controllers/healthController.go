package controllers

import "github.com/gin-gonic/gin"

// HealthCheck responds with a simple message indicating the server is up
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is up and running!",
	})
}
