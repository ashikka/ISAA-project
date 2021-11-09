package models

import (
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Id      string `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
