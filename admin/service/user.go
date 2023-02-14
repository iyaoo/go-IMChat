package service

import (
	"time"

	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/common/gorm"
)

func CreateUser() {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	user := models.User{}
	users := []models.User{}
	//e := db.AutoMigrate(&user).Error()
	//fmt.Println(e)

	user.Name = "Siri"
	//user.PassWord = "123456"
	user.Gender = "0"
	user.LoginTime = time.Now()
	user.LogoutTime = time.Now()
	err = db.Create(&user).Error
	if err != nil {
		slog.Errorf("create field err:%s", err)
	}
	err = db.Find(&users).Error
	if err != nil {
		slog.Errorf("select user err:%s", err)
	}
	slog.Info(users)
}
