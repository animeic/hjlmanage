package controllers

import (
	"animeic-gin/app/request"
	"animeic-gin/app/response"
	"animeic-gin/app/services"
	"animeic-gin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var AC = new(AuthController)

// 不需要鉴权
func (ac *AuthController) Register(c *gin.Context) {
	// 表单验证
	var form request.Register
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 验证用户名是否被注册
	// 注册时设置一个有序集合，username为member/分数为用户id user:username,user:email
	as := new(services.AuthService)
	score, _ := as.IsRegByName(form.Username)
	regUid := int64(score)
	if regUid > 0 {
		response.Fail(c, response.ExistUserErr)
		return
	}

	// 验证验证码是否正确
	isCode := as.IsCode(form.Etype, form.Email, form.Ecode)
	if !isCode {
		response.Fail(c, response.EmailCodeErr)
		return
	}
	// 用户信息写入数据库
	uid := as.RegisterUser(form)
	if uid < 1 {
		response.Fail(c, response.RegisterUserErr)
		return
	}
	user, _ := new(services.UserService).GetUserInfoById(uid)

	// 注册成功返回token
	jwt := new(services.JwtService)
	tokenData, _, err := jwt.CreateToken(services.AppGuardName, user)
	if err != nil {
		response.Fail(c, response.GenrateTokenError)
		return
	}

	response.Success(c, response.RegisterSuccess, tokenData)
}

// 根据用户名登录
func (ac *AuthController) Login(c *gin.Context) {
	// 接收用户名和密码
	var form request.LoginByName
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 验证用户名或者密码是否注册
	as := new(services.AuthService)
	score, err := as.IsRegByName(form.Username)
	uid := int64(score)
	// 查找不存在的key，返回redis nil,0
	if err != nil || uid == 0 {
		// response.Fail(c, response.DataNotFundErr)
		response.Fail(c, response.IsNotRegistUserErr)
		return
	}

	// 验证密码是否正确
	us := new(services.UserService)
	user, err := us.GetUserInfoByName(form.Username)
	if err != nil || user == nil {
		response.Fail(c, response.DataNotFundErr)
		return
	}
	if user.Password != utils.MD5([]byte(form.Password)) {
		response.Fail(c, response.LoginPasswordErr)
		return
	}
	// 登录返回token
	jwt := new(services.JwtService)
	tokenData, _, err := jwt.CreateToken(services.AppGuardName, user)
	if err != nil {
		response.Fail(c, response.GenrateTokenError)
		return
	}

	response.Success(c, response.LoginSuccess, tokenData)
}

// 发送邮箱验证码
func (ac *AuthController) SendCode(c *gin.Context) {
	var form request.SendCode
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	fmt.Println(form)
	as := new(services.AuthService)
	score, _ := as.IsRegByEmail(form.Email)
	regUid := int64(score)
	var type_text string
	switch form.Etype {
	case "register":

		// 验证邮箱是否被注册
		if regUid > 0 {
			response.Fail(c, response.ExistEmailErr)
			return
		}
		type_text = "注册账号"
	case "repass":
		if regUid == 0 {
			response.Fail(c, response.IsNotRegistUserErr)
			return
		}
		type_text = "重置密码"
	// 参数错误
	default:
		response.Fail(c, response.ParamsErr)
		return
	}

	// 邮箱模板
	// 发送目标邮箱验证码
	mailTo := []string{
		form.Email,
	}
	subject := "animeic.tech"
	code := utils.MakeEmailCode()
	body := fmt.Sprintf(`<p style="font-size:18px;">您正在<a href="http://127.0.0.1:3000">animeic.tech</a>%s，验证码是：
	<span style="color:blue;font-size:24px;">%s</span></p><p style="font-size:18x;">5分钟内有效，请妥善保管。</p>`, type_text, code)
	err = new(services.AuthService).SendMail(mailTo, subject, body)
	if err != nil {
		response.Fail(c, response.SendEmailCodeErr)
		return
	}
	// 发送成功生成缓存用于验证
	// mailcode:register:animeic@163.com
	// mailcode:repass:animeic@163.com
	rd := new(services.RecDb)
	err = rd.SetEmailCode(form.Etype, mailTo[0], code)
	if err != nil {
		response.Fail(c, response.SendEmailCodeErr)
		return
	}
	response.Success(c, response.SendEmailSuccess, nil)

}

// 重置密码
func (uc *AuthController) Repass(c *gin.Context) {
	var form request.Repass
	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	// 验证验证码
	as := new(services.AuthService)
	isCode := as.IsCode(form.Etype, form.Email, form.Ecode)
	if !isCode {
		response.Fail(c, response.EmailCodeErr)
		return
	}
	// 重置密码
	isrepass := new(services.AuthService).RepssUser(form)
	if !isrepass {
		response.Fail(c, response.RepssErr)
		return
	}
	response.Success(c, response.RepassSuccess, isrepass)
}

func (ac *AuthController) Print(c *gin.Context) {
	response.Success(c, "hello", nil)
}
