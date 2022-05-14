package bootstrap

import (
	"context"
	"fmt"

	"animeic-gin/global"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInit() {
	mgcf := global.App.Config.Mongo
	// 设置mongoDB客户端连接信息
	uri := fmt.Sprintf("mongodb://%s:%d", mgcf.DbName, mgcf.Port)
	cliopts := options.Client().ApplyURI(uri)
	var err error
	mgcli, err := mongo.Connect(context.TODO(), cliopts)
	if err != nil {
		global.App.Log.Info(err.Error())
		return
	}
	err = mgcli.Ping(context.TODO(), nil)
	if err != nil {
		global.App.Log.Info(err.Error())
		return
	}
	global.App.Mongo = mgcli
}
