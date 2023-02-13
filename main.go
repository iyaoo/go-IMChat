package main

import (
	"github.com/iyaoo/go-IMChat/admin/service"
)

func main() {
	service.CreateUser()
	//config.InitConfig()

	//str := config.App.Config.Settings.Database.Source
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(200, str)
	//})
	//r.Run()
}
