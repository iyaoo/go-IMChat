package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/common/apis"
)

type User struct {
	apis.Api
}

func (e *User) GetUserList(c *gin.Context) {
	data, err := service.SelectUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "not found data",
		})
	}
	db, err := e.GetOrm(c)
	if db == nil {
		slog.Error("get orm err", err)
	}
	slog.Info(db)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  data,
	})
}
