package dto

import (
	"github.com/iyaoo/go-IMChat/common/models"
)

// UserControl 用户增、改使用结构体
type UserControl struct {
	ID   int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:自增主键id"`
	Name string `json:"name" gorm:"column:name;not null;comment:用户姓名"`
	//PassWord      string    `json:"password" gorm:"column:password;not null;comment:用户密码"`
	Gender        string `json:"gender" gorm:"column:gender"`
	Phone         string `json:"phone" gorm:"column:phone"`
	Email         string `json:"email" gorm:"column:email"`
	Identity      string `json:"identity" gorm:"column:identity"`
	ClientIP      string `json:"client_ip" gorm:"column:client_ip"`
	ClientPost    string `json:"client_post" gorm:"column:client_post"`
	IsLogin       bool   `json:"is_login" gorm:"column:is_login"`
	HeartBeatTime string `json:"heartbeat_time" gorm:"column:heartbeat_time"`
	DeviceInfo    string `json:"device_info" gorm:"column:device_info"`
	//LoginTime     time.Time `json:"login_time" gorm:"column:login_time"`
	//LogoutTime    time.Time `json:"logout_time" gorm:"column:logout_time"`
	models.GormTime
}

// TableName 获取表名
func (table *UserControl) TableName() string {
	return "user"
}

// UserById 获取单个或者删除的结构体
type UserById struct {
	UserID int `uri:"id"`
}

func (u *UserById) GetId() interface{} {
	return u.UserID
}
