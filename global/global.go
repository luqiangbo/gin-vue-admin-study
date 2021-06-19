package global

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB    *gorm.DB
	GVA_REDIS *redis.Client
	GVA_VP    *viper.Viper
	GVA_LOG   *zap.Logger
	//GVA_Timer timer.Timer = timer.NewTimerTask()
)
