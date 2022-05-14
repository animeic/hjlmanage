package routes

import (
	"animeic-gin/app/controllers"
	"animeic-gin/app/middleware"
	"animeic-gin/app/services"
	"animeic-gin/app/ws"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	// api 跨域必须放在路由前面
	app.Use(middleware.Cors(), middleware.CustomRecovery())

	// 用户
	user := app.Group("/user")
	user.Use(middleware.Auth(services.AppGuardName))
	user.POST("/:id", controllers.UC.GetUserInfo)

	auth := app.Group("/auth")
	auth.POST("/register", controllers.AC.Register)
	auth.POST("/login", controllers.AC.Login)
	auth.POST("/sendCode", controllers.AC.SendCode)
	auth.POST("/repass", controllers.AC.Repass)
	auth.POST("/print", controllers.AC.Print)

	app.GET("/ws", ws.Handler)

	// auth := api.Group("/auth").Use(middleware.Cors(), middleware.JWTAuth(services.AppGuardName)) // iroutes

}
