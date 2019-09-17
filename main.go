package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", json)
	r.GET("/asciiJSON", asciiJSON)
	r.GET("/pureJson", pureJson)

	_ = r.Run(":8080")
}

func pureJson(c *gin.Context) {
	c.PureJSON(200, gin.H{"123": "<script>666测试</script>"})
}

func asciiJSON(c *gin.Context) {
	c.AsciiJSON(200, gin.H{"123": "<script>666测试</script>"})
}

func json(c *gin.Context) {
	c.JSON(200, gin.H{"message": "<script>666测试</script>"})
}
