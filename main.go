package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	// This is a placeholder for the main function.
	// In a real application, you would initialize your server, database connections, etc. here.
	server.Run(":8080") // Start the server on port 8080
}
