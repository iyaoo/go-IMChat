package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/common/apis"
)

type User struct {
	apis.Api
}

// GetUser
func (e *User) GetUser(c *gin.Context) {
	list := make([]models.User, 0)
	serviceUser := service.User{}
	err := serviceUser.Getuser(&list)
	if err != nil {
		slog.Info(err)
	}
	e.OK(c, list, "查询成功")
}

// GetUserList 查询所有User
func (e *User) GetUserList(c *gin.Context) {
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
