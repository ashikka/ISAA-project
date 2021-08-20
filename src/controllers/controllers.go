package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type blogPost struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

var blogPosts = []blogPost{
	{ID: "1", Title: "Blue Train", Author: "John Coltrane", Content: "Something"},
	{ID: "2", Title: "Jeru", Author: "Gerry Mulligan", Content: "Something"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Content: "Something"},
}

func GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, blogPosts)
	}
}

func PostBlogs() gin.HandlerFunc {
	var newBlog blogPost

	return func(c *gin.Context) {

		if err := c.BindJSON(&newBlog); err != nil {
			return
		}

		blogPosts = append(blogPosts, newBlog)
		c.IndentedJSON(http.StatusCreated, newBlog)
	}
}
