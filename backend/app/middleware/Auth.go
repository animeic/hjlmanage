package middleware

import (
	"animeic-gin/app/response"
	"animeic-gin/app/services"
	"animeic-gin/global"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		log.Println(tokenStr)
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// 截取"bearer "
		tokenStr = tokenStr[len(services.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		if err != nil || new(services.JwtService).IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// log.Println(claims.Id)
		// if !new(services.).IsExpires(claims.Id) {
		// 	response.TokenFail(c)
		// 	c.Abort()
		// 	return
		// }

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				user, err := new(services.JwtService).GetUserInfo(GuardName, claims.Id)
				if err != nil {
					global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := new(services.JwtService).CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = new(services.JwtService).JoinBlackList(token)
				}
			}
		}

		// 放在context中，方便登录后知道谁访问服务
		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
