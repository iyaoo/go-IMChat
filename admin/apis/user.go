package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iyaoo/go-IMChat/admin/service"
)

func User(c *gin.Context) {
	data, err := service.SelectUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "not found data",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  data,
	})
}
