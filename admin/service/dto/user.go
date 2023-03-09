package dto

import (
	"github.com/CoolBank/coinhub-base/logger"
	"github.com/gin-gonic/gin"
)

// UserControl 用户增、改使用结构体
type UserControl struct {
	Name     string `json:"name" gorm:"column:name;not null;comment:用户姓名"`
	PassWord string `json:"password" gorm:"column:password;not null;comment:用户密码"`
	Gender   string `json:"gender" gorm:"column:gender"`
	Phone    string `json:"phone" gorm:"column:phone"`
	Email    string `json:"email" gorm:"column:email"`
	Identity string `json:"identity" gorm:"column:identity"`
}

// TableName 获取表名
func (table *UserControl) TableName() string {
	return "user"
}

// Bind 映射上下文中的结构体数据
func (s *UserControl) Bind(ctx *gin.Context) error {
	err := ctx.Bind(s)
	if err != nil {
		logger.Debugf("Bind error: %s", err.Error())
		return err
	}
	return err
}

// UserById 获取单个或者删除的结构体
type UserById struct {
	UserID int `uri:"id"`
}

// GetId
func (s *UserById) GetId() interface{} {
	return s.UserID
}

// Bind
func (s *UserById) Bind(ctx *gin.Context) error {
	err := ctx.ShouldBindUri(s)
	if err != nil {
		logger.Debugf("ShouldBindUri error: %s", err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		logger.Debugf("ShouldBind error: %s", err.Error())
	}
	return err
}
