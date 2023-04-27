package global

import (
	"github.com/casbin/casbin/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"learn-gin/config"
)

var (
	CONFIG         config.Config
	DB             *gorm.DB
	Logger         *zap.SugaredLogger
	CasbinEnforcer *casbin.Enforcer
)
