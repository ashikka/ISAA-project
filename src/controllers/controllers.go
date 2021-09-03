package controllers

import (
	"ISAA-project/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetBlogs() gin.HandlerFunc {
	var blogPosts []models.BlogPost

	models.DB.Find(&blogPosts)
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

		models.DB.Create(&newBlog)
		c.IndentedJSON(http.StatusCreated, newBlog)
	}
}

func GetBlogByID() gin.HandlerFunc {
	var blogPosts []models.BlogPost

	return func(c *gin.Context) {
		id := c.Param("id")

		models.DB.First(&blogPosts, id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
	}
}

func DeleteBlogById() gin.HandlerFunc {
	var blogPost models.BlogPost

	return func(c *gin.Context) {
		id := c.Param("id")
		models.DB.First(&blogPost, id)
		models.DB.Delete(&blogPost)

		var blogPosts []models.BlogPost
		models.DB.Find(&blogPosts)

		c.IndentedJSON(http.StatusNotFound, blogPosts)

	}
}
