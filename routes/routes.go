package routes

import (
	"cypresslabs.org/backend/controllers"
	"github.com/gin-gonic/gin"
)

// Setup configures the API routes
func Setup(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
}
