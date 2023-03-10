package dto

import (
	"encoding/json"
	"time"

	"github.com/CoolBank/coinhub-base/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iyaoo/go-IMChat/admin/models"
)

type UserControl struct {
	UserID     int       `json:"id" uri:"id"`
	Name       string    `json:"name"`
	PassWord   string    `json:"password"`
	Gender     string    `json:"gender"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Identity   string    `json:"identity"`
	LoginTime  time.Time `json:"login_time"`
	LogoutTime time.Time `json:"logout_time"`
}

// Bind
func (s *UserControl) Bind(ctx *gin.Context) error {
	err := ctx.ShouldBindUri(s)
	if err != nil {
		logger.Debugf("ShouldBindUri error: %s", err.Error())
		return err
	}
	err = ctx.ShouldBindBodyWith(s, binding.JSON)
	if err != nil {
		logger.Debugf("ShouldBind error: %s", err.Error())
	}
	var jsonStr []byte
	jsonStr, err = json.Marshal(s)
	if err != nil {
		logger.Debugf("ShouldBind error: %s", err.Error())
	}
	ctx.Set("body", string(jsonStr))
	return err
}

// Generate 结构体数据转化 从 UserControl 至 models.User 对应的模型
func (s *UserControl) Generate() (*models.User, error) {
	return &models.User{
		ID:         s.UserID,
		Name:       s.Name,
		PassWord:   s.PassWord,
		Gender:     s.Gender,
		Phone:      s.Phone,
		Email:      s.Email,
		Identity:   s.Identity,
		LoginTime:  time.Now(),
		LogoutTime: time.Now(),
	}, nil
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
