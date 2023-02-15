package service

import (
	"time"

	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/admin/service/dto"
	"github.com/iyaoo/go-IMChat/common/gorm"
	"github.com/iyaoo/reusable-lib/tools"
)

// CreateUser
func CreateUser() {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	user := models.User{}
	users := []models.User{}
	//e := db.AutoMigrate(&user).Error()
	//fmt.Println(e)

	user.Name = "James"
	user.PassWord = "123456"
	user.Gender = "1"
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	//err = db.Create(&user).Error
	//if err != nil {
	//	slog.Errorf("create field err:%s", err)
	//}
	err = db.Find(&users).Error
	if err != nil {
		slog.Errorf("select user err:%s", err)
	}
	tools.PrintJSON(users)
}

// InsertUser
func InsertUser() {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	user := dto.UserControl{}
	users := []models.User{}
	user.Name = "Lisa"
	user.Gender = "1"
	if err = db.Create(&user).Error; err != nil {
		slog.Errorf("insert user failed:%s", err)
	}
	if err = db.Find(&users).Error; err != nil {
		slog.Errorf("select users failed:%s", err)
	}
	tools.PrintJSON(users)
}

// DeleteUser
func DeleteUser() {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	model := models.User{}
	err = db.Where("id>=?", 6).Delete(&model).Error
	if err != nil {
		slog.Error(err)
	}
	users := []models.User{}
	var count int64
	db.Find(&users).Count(&count)
	slog.Info(count)
}
func SelectUser() ([]*models.User, error) {
	data := make([]*models.User, 0)
	var count int64
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	err = db.Find(&data).Count(&count).Error
	if err != nil {
		slog.Errorf("select user failed:%s", err)
	}
	return data, nil
}
