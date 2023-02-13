package gorm

import (
	"github.com/iyaoo/go-IMChat/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() (*gorm.DB, error) {
	config.InitConfig()
	db, err := gorm.Open(mysql.Open(config.App.Config.Settings.Database.Source))
	if err != nil {
		return db, err
	}
	return db, nil
}
