package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserNoRouter)
	routerCheckRole = append(routerCheckRole, registerUserRouter)
}

// /api/v1
func registerUserRouter(v1 *gin.RouterGroup) {
	api := &apis.User{}
	r := v1.Group("/app")
	{
		r.GET("/user", api.GetUser)
		r.GET("/user/:id", api.GetUserByID)
		r.POST("/user", api.InsertUser)
		r.PUT("/user/:id", api.UpdateUser)
		r.DELETE("/user/:id", api.DeleteUser)
	}
}

// /v1
func registerUserNoRouter(v1 *gin.RouterGroup) {
	api := &apis.User{}
	r := v1.Group("/app").Use()
	{
		r.GET("/user", api.GetUser)
		r.GET("/user/:id", api.GetUserByID)
		r.POST("/user", api.InsertUser)
	}
}
