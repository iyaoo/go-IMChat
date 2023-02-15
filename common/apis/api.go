package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/iyaoo/reusable-lib/utils"
	"gorm.io/gorm"
)

type Api struct {
}

// GetOrm 获取Orm DB
func (e *Api) GetOrm(c *gin.Context) (*gorm.DB, error) {
	db, err := utils.GetOrm(c)
	if err != nil {
		slog.Error(c, http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}
