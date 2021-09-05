package main

import (
	"ISAA-project/src/controllers"
	"log"
	"ISAA-project/src/models"
	"ISAA-project/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {

    models.ConnectDB()
	router := gin.Default()
	router.Run(utils.GoDotEnvVariable("BACKEND_URL"))
	router.GET("/hi", controllers.Hello())
	router.GET("/blogs", controllers.GetBlogs())
	log.Fatal("hello")
	router.GET("/blogs/:id", controllers.GetBlogByID())
	router.POST("/blogs", controllers.PostBlogs())
	router.DELETE("/blogs/:id", controllers.DeleteBlogById())

}