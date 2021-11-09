package controllers

import (
	"ISAA-project/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "hello")
	}
}

func GetBlogs() gin.HandlerFunc {

	return func(c *gin.Context) {
		var blogPosts []models.BlogPost

		models.DB.Find(&blogPosts)
		c.IndentedJSON(http.StatusOK, blogPosts)
	}
}

func PostBlogs() gin.HandlerFunc {

	return func(c *gin.Context) {
		var newBlog models.BlogPost

		if err := c.BindJSON(&newBlog); err != nil {
			return
		}

		models.DB.Create(&newBlog)
		c.IndentedJSON(http.StatusCreated, newBlog)
	}
}

func GetBlogByID() gin.HandlerFunc {

	return func(c *gin.Context) {
		var blogPosts []models.BlogPost
		var blogPost models.BlogPost

		id := c.Param("id")

		result := models.DB.First(&blogPosts, id)
		if result.Error != nil {

			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, blogPost)
	}
}

func DeleteBlogById() gin.HandlerFunc {

	return func(c *gin.Context) {
		var blogPosts []models.BlogPost
		var blogPost models.BlogPost

		id := c.Param("id")
		result := models.DB.First(&blogPosts, id)
		if result.Error != nil {

			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Blog not found"})
			return
		}

		models.DB.Delete(&blogPost)

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})

	}
}
