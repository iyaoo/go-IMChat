package service

import (
	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/common/gorm"
	"github.com/iyaoo/go-IMChat/common/service"
)

type User struct {
	service.Service
}

// GetUser 获取User信息
func (e *User) GetUser(list *[]models.User) error {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	err = db.Find(list).Error
	if err != nil {
		slog.Errorf("get user failed:%s", err)
		return err
	}
	return nil
}
func (e *User) RemoveUser() error {
	return nil
}
