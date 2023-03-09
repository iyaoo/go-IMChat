package service

import (
	"errors"

	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service/dto"
	"github.com/iyaoo/go-IMChat/common/gormconn"
	"github.com/iyaoo/go-IMChat/common/service"
	"gorm.io/gorm"
)

type User struct {
	service.Service
}

// GetUser 获取User信息
func (e *User) GetUser(list *[]models.User) error {
	db, err := gormconn.InitGorm()
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
func (e *User) GetUserByID(d *dto.UserById, data *models.User) error {
	var err error
	var model models.User
	db, err := gormconn.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
		return err
	}

	db = db.Model(*&model).First(data, d.UserID)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: ", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
func (e *User) Insertuser(d *dto.UserControl) error {
	var err error
	var model models.User

	db, err := gormconn.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
		return err
	}

	err = db.Model(&model).Create(d).Error
	return err

}

// RemoveUser 根据id删除用户
func (e *User) RemoveUser(d *dto.UserById) error {
	var err error
	var data models.User
	db, err := gormconn.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	de := db.Model(&data).Unscoped().Delete(&data, d.GetId())
	if de.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete err:", err)
		return err
	}
	if de.RowsAffected == 0 {
		err = errors.New("无权删除数据")
		return err
	}
	return nil
}
