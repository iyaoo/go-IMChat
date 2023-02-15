package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/apis"
)

func UserRouter() *gin.Engine {
	api := &apis.User{}
	r := gin.Default()
	r.GET("/getUser", api.GetUserList)
	return r
}
