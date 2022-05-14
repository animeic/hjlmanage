package response

import (
	"animeic-gin/global"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrorCode int         `json:"code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ErrorCode: 0,
		Data:      data,
		Message:   message,
	})
}
func Fail(c *gin.Context, customErr AnErr) {
	c.JSON(http.StatusOK, Response{
		ErrorCode: customErr.Code,
		Data:      nil,
		Message:   customErr.Message,
	})
}

func ValidateFrom(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		ErrorCode: ValdateFormErrCode,
		Data:      nil,
		Message:   message,
	})
}
func TokenFail(c *gin.Context) {
	Fail(c, TokenError)
}

func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	// 非生产环境显示具体错误信息
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		nil,
		msg,
	})
	c.Abort()
}
