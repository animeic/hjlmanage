package services

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"animeic-gin/app/request"
	"animeic-gin/global"
	"animeic-gin/utils"

	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
)

type AuthService struct{}

const (
	//KCEWJSRYLYOVAOBD
	USER string = "animeic_vps@163.com"
	PASS string = "KCEWJSRYLYOVAOBD"
	HOST string = "smtp.163.com"
	PORT string = "465"
)

func (as *AuthService) SendMail(mailTo []string, subject string, body string) error {
	mailConfig := map[string]string{
		"user": USER,
		"pass": PASS,
		"host": HOST,
		"port": PORT,
	}
	port, _ := strconv.Atoi(mailConfig["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConfig["user"], "animeic官方"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConfig["host"], port, mailConfig["user"], mailConfig["pass"])

	err := d.DialAndSend(m)
	return err
}

// 是否注册 用户名
func (as *AuthService) IsRegByName(username string) (score float64, err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	score, err = rdc.ZScore(cxt, z_name_id, username).Result()
	if err != nil {
		// log.Println("username出错时zset分数是：-", score) //0
		global.App.Log.Error(err.Error())
		return
	}
	return

}

// 是否注册 邮箱
func (as *AuthService) IsRegByEmail(email string) (score float64, err error) {
	rdc := global.App.Redis
	cxt := context.Background()
	score, err = rdc.ZScore(cxt, z_email_id, email).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	return
}

// 验证码是否正确
func (as *AuthService) IsCode(etype string, toemail string, ecode string) bool {
	rdc := global.App.Redis
	cxt := context.Background()
	key := fmt.Sprintf("ecode:%s:%s", etype, toemail)
	code1, err := rdc.Get(cxt, key).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return false
	}

	if code1 == ecode {
		return true
	}
	return false
}

// 注册用户 返回用户id
func (as *AuthService) RegisterUser(user request.Register) int64 {
	rdc := global.App.Redis
	cxt := context.Background()
	// 获取用户id
	// id生成器
	uid := new(RecDb).GenrateId("user")
	if uid < 1 {
		return 0
	}
	// 设置用户信息
	key := fmt.Sprintf("user:%d", uid)
	password := utils.MD5([]byte(user.Password))
	userinfo := map[string]interface{}{
		"id":       uid,
		"email":    user.Email,
		"username": user.Username,
		"password": password,
		"created":  time.Now().Unix(),
		"updated":  time.Now().Unix(),
	}

	_, err := rdc.HSet(cxt, key, userinfo).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return 0
	}
	zukey := "user:username"
	zekey := "user:email"
	rdc.ZAdd(cxt, zukey, &redis.Z{
		Score:  float64(uid),
		Member: user.Username,
	})
	rdc.ZAdd(cxt, zekey, &redis.Z{
		Score:  float64(uid),
		Member: user.Email,
	})

	return uid
}

// 注册用户 返回用户id
func (as *AuthService) RepssUser(form request.Repass) bool {
	rdc := global.App.Redis
	cxt := context.Background()
	// 获取用户id
	user1, err := new(UserService).GetUserInfoByEmail(form.Email)
	if err != nil {
		global.App.Log.Error(err.Error())
		return false
	}
	// 设置用户信息
	key := fmt.Sprintf("user:%s", user1.Id)
	password := utils.MD5([]byte(form.Password))
	_, err = rdc.HSet(cxt, key, "password", password, "updated", time.Now().Unix()).Result()
	if err != nil {
		global.App.Log.Error(err.Error())
		return false
	}

	return true
}
