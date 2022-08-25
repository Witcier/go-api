package global

import (
	"witcier/go-api/config"

	"golang.org/x/sync/singleflight"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	DBList             map[string]*gorm.DB
	Redis              *redis.Client
	Viper              *viper.Viper
	Log                *zap.Logger
	Config             config.Server
	ConcurrencyControl = &singleflight.Group{}
)

func GetGlobalDBByName(dbName string) *gorm.DB {
	return DBList[dbName]
}
