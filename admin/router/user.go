package router

import "github.com/gin-gonic/gin"

func User() {
	r := gin.Default()
	r.GET("/getUser")
}
