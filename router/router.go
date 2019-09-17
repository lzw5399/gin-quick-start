package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	group := r.Group("/user", customHandler)

	{
		// user/login
		group.Any("/login", loginHandler)
	}
	r.Run()
}

func loginHandler(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		fmt.Println("====== Only Bind By Query String ======")
		fmt.Println(person.Name)
		fmt.Println(person.Address)
	}
	c.String(200, "Success")
}

func customHandler(c *gin.Context) {
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}
