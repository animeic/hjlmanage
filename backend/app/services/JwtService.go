package services

import (
	"animeic-gin/global"
	"animeic-gin/utils"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type JwtService struct{}

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (js *JwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, token *jwt.Token, err error) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl,
				Id:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))
	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenType,
	}
	return
}

// 获取黑名单
func (js *JwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

func (js *JwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	err = global.App.Redis.SetNX(context.Background(), js.getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

func (jwtService *JwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, _ := global.App.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true

}

func (js *JwtService) GetUserInfo(GuardName string, id string) (user JwtUser, err error) {
	switch GuardName {
	case AppGuardName:
		return new(UserService).GetUserInfo(id)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
