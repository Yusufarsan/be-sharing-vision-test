package models

import (
	"time"
)

type Article struct {
	Id          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null" json:"status"`
}

func (Article) TableName() string {
	return "posts"
}