package main

import (
	"cypresslabs.org/backend/config"
	"cypresslabs.org/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize configuration
	config.Init()

	// Setup routes
	routes.Setup(router)

	// Start the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
