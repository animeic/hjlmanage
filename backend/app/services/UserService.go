package services

import (
	"animeic-gin/app/models"
	"animeic-gin/global"
	"context"
	"fmt"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type UserService struct{}

// jwt中获取用户的方法
func (us *UserService) GetUserInfo(id string) (user *models.User, err error) {
	intId, _ := strconv.ParseInt(id, 10, 64)
	user, err = getUserInfoById(intId)
	return

}

// 根据用户账号获取用户信息
func (us *UserService) GetUserInfoByName(username string) (user *models.User, err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	key := "user:username"

	// 在有序集合中查找id
	score, err := rdc.ZScore(cxt, key, username).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	id := int64(score)
	// 根据id获取用户信息
	user, err = getUserInfoById(id)
	return
}

// 根据邮箱获取用户信息
func (us *UserService) GetUserInfoByEmail(email string) (user *models.User, err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	key := "user:email"

	// 在有序集合中查找id
	score, err := rdc.ZScore(cxt, key, email).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	id := int64(score)
	// 根据id获取用户信息
	user, err = getUserInfoById(id)
	return
}
func (us *UserService) GetUserInfoById(id int64) (user *models.User, err error) {
	user, err = getUserInfoById(id)
	return
}

// 根据id获取用户信息
func getUserInfoById(id int64) (user *models.User, err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	key := fmt.Sprintf("user:%d", id)
	uf1, err := rdc.HGetAll(cxt, key).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	// todo hashmap key和value都是字符串，如何解码 id int64
	mapstructure.Decode(uf1, &user)
	return
}
