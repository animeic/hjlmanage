package bootstrap

import (
	"animeic-gin/global"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() {
	// 获取redis配置
	redisConfig := global.App.Config.Redis

	// 创建redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed,err:", zap.Any("err", err))
		return
	}
	// 挂载到全局变量上
	global.App.Redis = client

}
