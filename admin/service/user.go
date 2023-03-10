package service

import (
	"errors"

	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service/dto"
	"github.com/iyaoo/go-IMChat/common/gormconn"
	"github.com/iyaoo/go-IMChat/common/service"
	"github.com/iyaoo/reusable-lib/tools/logger"
	"gorm.io/gorm"
)

type User struct {
	service.Service
}

// GetUser 获取User信息
func (e *User) GetUser(list *[]models.User) error {
	db, err := gormconn.InitGorm()
	if err != nil {
		logger.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	err = db.Find(list).Error
	if err != nil {
		logger.Errorf("get user failed:%s", err)
		return err
	}
	return nil
}
func (e *User) GetUserByID(d *dto.UserById, data *models.User) error {
	var err error
	var model models.User
	orm, err := gormconn.InitGorm()
	if err != nil {
		logger.Errorf("gorm connect mysql failed:%s", err)
		return err
	}

	db := orm.Model(*&model).First(data, d.UserID)
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

// Insertuser
func (e *User) InsertUser(data *models.User) error {
	var err error
	var model models.User
	db, err := gormconn.InitGorm()
	if err != nil {
		logger.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	err = db.Model(&model).Create(data).Error
	if err != nil {
		logger.Debugf("db error: %s", err)
		return err
	}
	return nil
}

// UpdateUser 修改用户信息
func (e *User) UpdateUser(data *models.User) error {
	var err error
	var model models.User
	orm, err := gormconn.InitGorm()
	if err != nil {
		logger.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	tx := orm.Begin()
	db := tx.Model(&model).Where(data.GetId()).Updates(data)
	if db.Error != nil {
		err = db.Error
		tx.Rollback()
		logger.Errorf("db error: %s", err)
		return err
	}

	tx.Commit()
	return nil
}

// RemoveUser 删除用户
func (e *User) RemoveUser(d *dto.UserById) error {
	var err error
	var model models.User

	orm, err := gormconn.InitGorm()
	if err != nil {
		logger.Errorf("gorm connect mysql failed:%s", err)
		return err
	}
	//永久删除
	//db := orm.Model(&model).Unscoped().Delete(&model, d.GetId())
	//软删除
	db := orm.Model(&model).Delete(&model, d.GetId())
	if db.Error != nil {
		err = db.Error
		logger.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
