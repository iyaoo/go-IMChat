package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/common/apis"
)

type User struct {
	apis.Api
}

// GetUser 获取user列表
func (e *User) GetUser(c *gin.Context) {
	list := make([]models.User, 0)
	serviceUser := service.User{}
	err := serviceUser.Getuser(&list)
	if err != nil {
		slog.Info(err)
	}
	e.OK(c, list, "查询成功")
}
