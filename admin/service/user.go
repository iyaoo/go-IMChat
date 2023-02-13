package service

import (
	"fmt"

	"github.com/gookit/slog"
	"github.com/iyaoo/go-IMChat/admin/models"
	"github.com/iyaoo/go-IMChat/common/gorm"
)

func CreateUser() {
	db, err := gorm.InitGorm()
	if err != nil {
		slog.Errorf("gorm connect mysql failed:%s", err)
	}
	e := db.AutoMigrate(&models.User{}).Error()
	fmt.Println(e)
}
