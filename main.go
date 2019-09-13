package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handleHome)
	r.GET("/json", handleJson)

	_ = r.Run(":8081")
}

func handleJson(c *gin.Context) {
	c.AsciiJSON(200, gin.H{"123": "测试"})
}

func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "666"})
}
