package initialize

import (
	"context"
	"witcier/go-api/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Log.Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
}
