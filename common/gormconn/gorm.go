package gormconn

import (
	"github.com/iyaoo/go-IMChat/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGorm 初始化gorm连接
func InitGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.App.Config.Settings.Database.Source))
	if err != nil {
		return db, err
	}
	return db, nil
}
