package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handleHome)

	_ = r.Run()
}

func handleHome(context *gin.Context) {
	context.JSON(200, gin.H{"message": "666"})
}
