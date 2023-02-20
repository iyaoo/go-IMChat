package apis

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
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
