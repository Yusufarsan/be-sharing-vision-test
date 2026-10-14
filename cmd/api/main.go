package main

import (
	"be-sharing-vision-test/internal/config"
	"be-sharing-vision-test/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inisialisasi Database
	config.ConnectDatabase()

	// 2. Auto Migrate (Membuat tabel 'posts' jika belum ada)
	err := config.DB.AutoMigrate(&models.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed.")

	// 3. Setup Gin Router
	r := gin.Default()

	// Tes rute sederhana
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 4. Jalankan Server
	r.Run(":8080")
}