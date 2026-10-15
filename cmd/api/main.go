package main

import (
	"be-sharing-vision-test/internal/config"
	"be-sharing-vision-test/internal/models"
	"be-sharing-vision-test/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed.")

	r := gin.Default()
	r.Use(cors.Default())

	// Inisialisasi routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}