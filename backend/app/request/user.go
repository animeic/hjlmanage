package request

// 验证获取用户信息
type Id struct {
	Id int64 `form:"id" json:"id" binding:"required,gte=0"`
}

func (id Id) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "账号不能为空",
	}
}

type UserName struct {
	Username string `form:"username" json:"username" binding:"required,username,min=4,max=16"`
}

func (username UserName) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"username.required": "账号不能为空",
	}
}

type Email struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

func (email Email) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"email.required": "账号不能为空！",
		"email.email":    "邮箱格式不正确！",
	}
}
