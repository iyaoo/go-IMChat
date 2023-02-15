package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/apis"
)

func UserRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getUser", apis.User)
	return r
}
