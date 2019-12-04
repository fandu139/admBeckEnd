package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Server Sudah Jalan",
		})
	})

	route.POST("/auth", func(context *gin.Context) {
		name := context.PostForm("name")
		message := context.PostForm("message")

		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"name":    name,
		})

		// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	route.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
