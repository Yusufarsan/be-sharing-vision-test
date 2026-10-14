package main

import (
	"log"
	"net/http"
	"os"

	"be-sharing-vision-test/internal/config"
	"be-sharing-vision-test/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	r := gin.Default()
	r.Use(config.CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	routes.RegisterArticleRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
