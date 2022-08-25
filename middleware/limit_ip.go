package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"
	"witcier/go-api/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LimitConfig struct {
	GenerationKey func(c *gin.Context) string
	CheckOrMark   func(key string, expire int, limit int) error
	Expire        int
	Limit         int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err,
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func DefaultGenerationKey(c *gin.Context) string {
	return "API_LIMIT" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	if global.Redis == nil {
		return err
	}

	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.Log.Error("limit", zap.Error(err))
	}

	return err
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Limit:         global.Config.System.IplimitTime,
		Expire:        global.Config.System.IplimitCount,
	}.LimitWithTime()
}

func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.Redis.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if count == 0 {
		pipe := global.Redis.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err := pipe.Exec(context.Background())

		return err
	} else {
		if times, err := global.Redis.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.Redis.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.Redis.Incr(context.Background(), key).Err()
			}
		}
	}
}
