package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name          string
	PassWord      string `gorm:"column:password"`
	Gender        string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string    `gorm:"column:client_ip"`
	ClientPost    string    `gorm:"column:client_post"`
	LoginTime     time.Time `gorm:"column:login_time"`
	LogoutTime    time.Time `gorm:"column:logout_time"`
	IsLogout      bool      `gorm:"column:is_logout"`
	HeartBeatTime string    `gorm:"column:heartbeat_time"`
	DeviceInfo    string    `gorm:"column:device_info"`
	//Created_at    string
	//Updated_at    string
	//Deleted_at    string
	gorm.Model
}

func (table *User) TableName() string {
	return "users"
}
