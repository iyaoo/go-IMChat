package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserNoRouter)
	routerCheckRole = append(routerCheckRole, registerUserRouter)
}

func registerUserRouter(v1 *gin.RouterGroup) {
	api := &apis.User{}
	{
		v1.GET("/getuser", api.GetUser)
	}
}
func registerUserNoRouter(v1 *gin.RouterGroup) {
	api := &apis.User{}
	{
		v1.GET("/user", api.GetUser)
	}
}
