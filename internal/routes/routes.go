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
		articleGroup.GET("/:p1/:p2", controllers.GetArticles)	
		articleGroup.GET("/:p1", controllers.GetArticleByID)
		articleGroup.PUT("/:id", controllers.UpdateArticle)
		articleGroup.DELETE("/:id", controllers.DeleteArticle)
	}
}