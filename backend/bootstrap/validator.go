package bootstrap

import (
	"animeic-gin/utils"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidator() {
	// 自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)
		_ = v.RegisterValidation("email", utils.ValidateEmail)
		_ = v.RegisterValidation("username", utils.ValidateUserName)
		_ = v.RegisterValidation("password", utils.ValidatePassword)

		// 注册自定义json tag函数
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 这里有空格号分割
			name := strings.SplitN(field.Tag.Get("json"), " ", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
