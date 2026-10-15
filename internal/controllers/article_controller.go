package controllers

import (
	"be-sharing-vision-test/internal/config"
	"be-sharing-vision-test/internal/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationError merepresentasikan satu error validasi per field
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// formatValidationErrors mengubah error dari Gin/validator menjadi array terstruktur
func formatValidationErrors(err error) []ValidationError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ValidationError, len(ve))
		for i, fe := range ve {
			out[i] = ValidationError{
				Field:   strings.ToLower(fe.Field()),
				Message: getErrorMessage(fe),
			}
		}
		return out
	}

	// Fallback untuk error non-validasi (misal JSON syntax error)
	return []ValidationError{
		{Field: "general", Message: err.Error()},
	}
}

// getErrorMessage menghasilkan pesan error yang human-readable berdasarkan tag validasi
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
	case "oneof":
		options := strings.Split(fe.Param(), " ")
		if len(options) > 1 {
			last := options[len(options)-1]
			rest := strings.Join(options[:len(options)-1], ", ")
			return fmt.Sprintf("%s must be either %s, or %s", fe.Field(), rest, last)
		}
		return fmt.Sprintf("%s must be one of: %s", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}

// 1. POST /article/
func CreateArticle(c *gin.Context) {
	var input models.ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": formatValidationErrors(err),
		})
		return
	}

	article := models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	}

	config.DB.Create(&article)
	c.JSON(http.StatusOK, gin.H{})
}

// 2. GET /article/list/:limit/:offset
func GetArticles(c *gin.Context) {
	limitStr := c.Param("p1")	// p1 disini adalah limit
	offsetStr := c.Param("p2")	// p2 disini adalah offset

	limit, errLimit := strconv.Atoi(limitStr)
	offset, errOffset := strconv.Atoi(offsetStr)

	if errLimit != nil || errOffset != nil || limit < 0 || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []ValidationError{
				{Field: "limit", Message: "Invalid limit or offset parameters"},
			},
		})
		return
	}

	var articles []models.Article
	config.DB.Limit(limit).Offset(offset).Find(&articles)

	c.JSON(http.StatusOK, articles)
}

// 3. GET /article/:id
func GetArticleByID(c *gin.Context) {
	id := c.Param("p1")	// p1 disini adalah id, bukan limit
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []ValidationError{
				{Field: "id", Message: "Article not found"},
			},
		})
		return
	}

	c.JSON(http.StatusOK, article)
}

// 4. PUT /article/:id
func UpdateArticle(c *gin.Context) {
	id := c.Param("p1")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []ValidationError{
				{Field: "id", Message: "Article not found"},
			},
		})
		return
	}

	var input models.ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": formatValidationErrors(err),
		})
		return
	}

	config.DB.Model(&article).Updates(models.Article{
		Title:    input.Title,
		Content:  input.Content,
		Category: input.Category,
		Status:   input.Status,
	})

	c.JSON(http.StatusOK, gin.H{})
}

// 5. DELETE /article/:id (Diimplementasikan sebagai Soft Delete ke status 'thrash')
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []ValidationError{
				{Field: "id", Message: "Article not found"},
			},
		})
		return
	}

	// Soft delete mengubah status, bukan menghapus record secara permanen
	config.DB.Model(&article).Update("status", "thrash")

	c.JSON(http.StatusOK, gin.H{"message": "Article moved to thrash"})
}