package apis

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/reusable-lib/sdk/pkg/response"
	"gorm.io/gorm"
)

type Api struct {
}

// GetOrm
func (e *Api) GetOrm(c *gin.Context) (*gorm.DB, error) {
	db, err := GetOrmConnect(c)
	if err != nil {
		slog.Error(c, http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

// GetOrmConnect 获取orm连接
func GetOrmConnect(c *gin.Context) (*gorm.DB, error) {
	idb, exist := c.Get("db")
	if !exist {
		return nil, errors.New("err1-db connect not exist")
	}
	switch idb.(type) {
	case *gorm.DB:
		//新增操作
		return idb.(*gorm.DB), nil
	default:
		return nil, errors.New("err2-db connect not exist")
	}
}

// Error 通常错误数据处理
func (e *Api) Error(c *gin.Context, code int, err error, msg string) {
	response.Error(c, code, err, msg)
}

// OK 通常成功数据处理
func (e *Api) OK(c *gin.Context, data interface{}, msg string) {
	response.OK(c, data, msg)
}

// PageOK分页数据处理
func (e *Api) PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(c, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e *Api) Custom(c *gin.Context, data gin.H) {
	response.Custum(c, data)
}
