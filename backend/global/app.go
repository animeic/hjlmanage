package global

import (
	"animeic-gin/app/models"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      models.Config
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
	Mongo       *mongo.Client
}

// 读取配置的入口
var App = new(Application)
