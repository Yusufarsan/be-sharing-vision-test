package models

type ArticleInput struct {
	Title    string `json:"title" binding:"required,min=20"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status" binding:"required,oneof=publish draft thrash"`
}