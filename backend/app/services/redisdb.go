package services

import (
	"animeic-gin/global"
	"context"
	"fmt"
	"time"
)

type RecDb struct {
	// rdc *redis.Client
	// cxt context.Context
}

// 设置邮箱验证码
func (rb *RecDb) SetEmailCode(etype string, toemail string, ecode string) (err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	key := fmt.Sprintf("ecode:%s:%s", etype, toemail)
	expiration := 5 * 60 * 26 * 60 * time.Second
	err = rdc.SetEX(cxt, key, ecode, expiration).Err()
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	return
}

// id生成

func (rb *RecDb) GenrateId(ktype string) int64 {
	rdc := global.App.Redis
	cxt := context.Background()
	// user:uid
	key := fmt.Sprintf("uid:%s", ktype)
	oid, err := rdc.BitPos(cxt, key, 0).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return 0
	}
	uid := oid + 1
	_, err1 := rdc.SetBit(cxt, key, oid, 1).Result()
	if err1 != nil {
		global.App.Log.Error(err.Error())
		return 0
	}
	return uid
}
