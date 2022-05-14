package response

type AnErr struct {
	Code    int
	Message string
}

var (
	ValdateFormErrCode int    = 1001
	WebSocketCode      int    = 1002
	RegisterSuccess    string = "注册成功！"
	SendEmailSuccess   string = "验证码发送成功！"
	GetUserInfoSuccess string = "获取用户信息成功！"
	LoginSuccess       string = "登录成功！"
	RepassSuccess      string = "重置密码成功！"
	ParamsErr          AnErr  = AnErr{Code: 1010, Message: "参数错误！"}
	DataNotFundErr     AnErr  = AnErr{Code: 2000, Message: "数据查找失败！"}
	RegisterError      AnErr  = AnErr{Code: 2001, Message: "注册失败！"}
	SendEmailCodeErr   AnErr  = AnErr{Code: 2002, Message: "发送邮件失败！"}
	EmailCodeErr       AnErr  = AnErr{Code: 2003, Message: "邮箱验证码错误！"}
	RegisterUserErr    AnErr  = AnErr{Code: 2004, Message: "用户信息注册失败！"}
	GenrateTokenError  AnErr  = AnErr{Code: 2005, Message: "令牌生成失败！"}
	GetUserInfoErr     AnErr  = AnErr{Code: 2006, Message: "获取用户信息失败！"}
	ExistUserErr       AnErr  = AnErr{Code: 2008, Message: "该用户已被注册！"}
	ExistEmailErr      AnErr  = AnErr{Code: 2009, Message: "该邮箱已被注册！"}
	IsNotRegistUserErr AnErr  = AnErr{Code: 2010, Message: "该用户未注册！"}
	LoginPasswordErr   AnErr  = AnErr{Code: 2011, Message: "密码错误！"}
	RepssErr           AnErr  = AnErr{Code: 2012, Message: "密码重置失败！"}

	TokenError AnErr = AnErr{Code: 2002, Message: "无效令牌！"}
)
