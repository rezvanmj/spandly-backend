package main

import (
	"log"
	"spandly-backend/internal/db"
	"spandly-backend/internal/route"
	"spandly-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db.Connect()

	// Run migrations
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	// Gin setup
	r := gin.Default()

	// Register routes
	route.RegisterUserRoutes(r)

	// Start server
	r.Run(":8080")
}
