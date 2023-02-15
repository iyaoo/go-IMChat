package service

import (
	"github.com/gookit/slog"
	"gorm.io/gorm"
)

type Service struct {
	Orm *gorm.DB
	Log *slog.Logger
}
