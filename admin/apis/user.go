package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/admin/service/dto"
	"github.com/iyaoo/go-IMChat/common/apis"
)

type User struct {
	apis.Api
}

// GetUser 获取user列表
func (e *User) GetUser(c *gin.Context) {
	list := make([]models.User, 0)
	serviceUser := service.User{}
	err := serviceUser.GetUser(&list)
	if err != nil {
		slog.Info(err)
	}
	e.OK(c, list, "查询成功")
}

// GetUserByID 根据ID获取用户信息
func (e *User) GetUserByID(c *gin.Context) {
	control := new(dto.UserById)

	err := control.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.User

	serviceUser := service.User{}
	err = serviceUser.GetUserByID(control, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(c, object, "查看成功")
}
