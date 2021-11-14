package main

import (
	"ISAA-project/src/controllers"
	"ISAA-project/src/models"
	"github.com/gin-contrib/cors"
	"ISAA-project/src/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

    models.ConnectDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"POST", "GET"},
        AllowHeaders:     []string{"Origin", "Content-type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))
	router.GET("/hi", controllers.Hello())
	router.GET("/blogs", controllers.GetBlogs())
	router.GET("/blogs/:id", controllers.GetBlogByID())
	router.POST("/blogs", controllers.PostBlogs())
	router.DELETE("/blogs/:id", controllers.DeleteBlogById())
	router.Run(utils.GoDotEnvVariable("BACKEND_URL"))
}