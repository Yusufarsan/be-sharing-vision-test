package routes

import (
	"be-sharing-vision-test/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterArticleRoutes(router *gin.Engine, db *gorm.DB) {
	articleController := controllers.NewArticleController(db)

	article := router.Group("/article")
	{
		article.POST("/", articleController.CreateArticle)
		article.GET("/:limit/:offset", articleController.GetArticles)
		article.GET("/:id", articleController.GetArticle)
		article.POST("/:id", articleController.UpdateArticle)
		article.PUT("/:id", articleController.UpdateArticle)
		article.PATCH("/:id", articleController.UpdateArticle)
		article.DELETE("/:id", articleController.DeleteArticle)
	}
}
