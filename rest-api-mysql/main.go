package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Gin in Docker!",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	} // Default listen on :8080
}
