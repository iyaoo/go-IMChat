package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/common/config"
)

func main() {
	//service.CreateUser()
	//service.InsertUser()
	service.DeleteUser()
	config.InitConfig()

	str := config.App.Config.Settings.Database.Source
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, str)
	})
	r.Run()
}
