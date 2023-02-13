package models

import "gorm.io/gorm"

type User struct {
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	CLientPost    string
	LoginTime     uint64
	LogoutTime    uint64
	IsLogout      bool
	HeartbeatTime uint64
	DeviceInfo    string
	gorm.Model
}

func (table *User) TableName() string {
	return "user"
}
