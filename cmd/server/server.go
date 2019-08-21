package main

import "github.com/gin-gonic/gin"

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test2",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", test)
	router.Run()
}