package models

import (
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
