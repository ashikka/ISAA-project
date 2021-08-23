package controllers

import (
	"ISAA-project/src/models"
	"ISAA-project/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = models.ConnectDB(utils.GoDotEnvVariable("DB_URL"))

func GetBlogs() gin.HandlerFunc {
	var blogPosts []models.BlogPost

	db.Find(&blogPosts)
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, blogPosts)
	}
}

func PostBlogs() gin.HandlerFunc {
	var newBlog models.BlogPost

	return func(c *gin.Context) {

		if err := c.BindJSON(&newBlog); err != nil {
			return
		}

		db.Create(&newBlog)
		c.IndentedJSON(http.StatusCreated, newBlog)
	}
}

func GetBlogByID() gin.HandlerFunc {
	var blogPosts []models.BlogPost

	return func(c *gin.Context) {
		id := c.Param("id")

		db.First(&blogPosts, id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
	}
}

func DeleteBlogById() gin.HandlerFunc {
	var blogPost models.BlogPost

	return func(c *gin.Context) {
		id := c.Param("id")
		db.First(&blogPost, id)
		db.Delete(&blogPost)

		var blogPosts []models.BlogPost
		db.Find(&blogPosts)

		c.IndentedJSON(http.StatusNotFound, blogPosts)

	}
}
