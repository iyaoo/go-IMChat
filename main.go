package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "dangdangdangdang~")
	})
	r.Run()
}
