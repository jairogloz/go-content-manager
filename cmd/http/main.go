package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/content", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello go!",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 by default
}
