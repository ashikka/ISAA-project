package models

import (
	"github.com/jinzhu/gorm"
)

type BlogPost struct {
	gorm.Model
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
