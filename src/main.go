package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func main() {
	router := gin.Default()
	router.GET("/blogs", getBlogs)

	router.Run(goDotEnvVariable("BACKEND_URL"))
}

func getBlogs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blogPosts)
}
