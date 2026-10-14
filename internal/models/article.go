package models

import "time"

type Article struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(200);not null" json:"title"`
	Content     string    `gorm:"column:content;type:text;not null" json:"content"`
	Category    string    `gorm:"column:category;type:varchar(100);not null" json:"category"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"column:status;type:varchar(100);not null" json:"status"`
}

func (Article) TableName() string {
	return "posts"
}
