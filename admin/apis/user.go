package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service"
	"github.com/iyaoo/go-IMChat/admin/service/dto"
	"github.com/iyaoo/go-IMChat/common/apis"
	"github.com/iyaoo/reusable-lib/tools/logger"
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

// InsertUser 添加用户
func (e *User) InsertUser(c *gin.Context) {
	control := new(dto.UserControl)
	err := control.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.Generate()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	serviceUser := service.User{}
	err = serviceUser.InsertUser(object)
	if err != nil {
		logger.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

// UpdateUser 更新用户信息
func (e *User) UpdateUser(c *gin.Context) {
	control := new(dto.UserControl)
	err := control.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, err := control.Generate()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	serviceUser := service.User{}
	err = serviceUser.UpdateUser(object)
	if err != nil {
		logger.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

// DeleteUser 删除用户
func (e *User) DeleteUser(c *gin.Context) {
	control := new(dto.UserById)
	err := control.Bind(c)
	if err != nil {
		logger.Errorf("Bind error: %s", err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	serviceUser := service.User{}
	err = serviceUser.RemoveUser(control)
	if err != nil {
		logger.Error(err)
		return
	}
	e.OK(c, control.GetId(), "删除成功")
}
