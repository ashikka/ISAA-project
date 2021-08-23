package main

import (
	"ISAA-project/src/controllers"
	"ISAA-project/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/blogs", controllers.GetBlogs())
	router.GET("/blogs/:id", controllers.GetBlogByID())
	router.POST("/blogs", controllers.PostBlogs())
	router.DELETE("/blogs/:id", controllers.DeleteBlogById())

	router.Run(utils.GoDotEnvVariable("BACKEND_URL"))
}
