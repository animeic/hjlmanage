package controllers

import (
	"animeic-gin/app/models"
	"animeic-gin/app/request"
	"animeic-gin/app/response"
	"animeic-gin/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var UC = new(UserController)

func (uc *UserController) GetUserInfo(c *gin.Context) {
	// 根据id，用户名，邮箱获取用户信息
	var form request.Id
	err := c.ShouldBindQuery(&form)
	if err != nil {
		response.ValidateFrom(c, request.GetErrorMsg(&form, err))
		return
	}
	var user *models.User

	if form.Id > 0 {
		user, err = new(services.UserService).GetUserInfoById(form.Id)
	}
	if err != nil {
		response.Fail(c, response.GetUserInfoErr)
		return
	}
	response.Success(c, response.GetUserInfoSuccess, user)
}
