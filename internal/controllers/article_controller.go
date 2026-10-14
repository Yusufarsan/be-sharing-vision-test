package controllers

import (
	"be-sharing-vision-test/internal/config"
	"be-sharing-vision-test/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 1. POST /article/
func CreateArticle(c *gin.Context) {
	var input models.ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	}

	config.DB.Create(&article)
	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully"})
}

// 2. GET /article/list/:limit/:offset
func GetArticles(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)

	if errLimit != nil || errOffset != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit or offset parameters"})
		return
	}

	var articles []models.Article
	config.DB.Limit(limit).Offset(offset).Find(&articles)

	c.JSON(http.StatusOK, articles)
}

// 3. GET /article/:id
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// 4. PUT /article/:id (Bisa juga POST/PATCH sesuai spesifikasi)
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	var input models.ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&article).Updates(models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Article updated successfully"})
}

// 5. DELETE /article/:id (Diimplementasikan sebagai Soft Delete ke status 'thrash')
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Soft delete mengubah status, bukan menghapus record secara permanen
	config.DB.Model(&article).Update("status", "thrash")

	c.JSON(http.StatusOK, gin.H{"message": "Article moved to thrash"})
}