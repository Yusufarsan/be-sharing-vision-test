package routes

import (
	"be-sharing-vision-test/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Grouping rute /article
	articleGroup := r.Group("/article")
	{
		articleGroup.POST("/", controllers.CreateArticle)
		articleGroup.GET("/list/:limit/:offset", controllers.GetArticles)	
		articleGroup.GET("/:id", controllers.GetArticleByID)
		articleGroup.PUT("/:id", controllers.UpdateArticle)
		// atau POST /:id untuk update, sesuai preferensi, di sini menggunakan standar REST PUT
		articleGroup.DELETE("/:id", controllers.DeleteArticle)
	}
}