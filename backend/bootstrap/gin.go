package bootstrap

import (
	"animeic-gin/app/ws"
	"animeic-gin/global"
	"animeic-gin/routes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	app := gin.Default()
	routes.InitRouter(app)
	gin.SetMode(global.App.Config.App.Mode)
	handleSignal(app)
	go ws.Manager.Start()
	app.Run(fmt.Sprintf("%s:%s", global.App.Config.App.AppUrl, global.App.Config.App.Port))

	// go1.8版本关闭服务

	// srv := &http.Server{
	// 	Addr:    fmt.Sprintf("%s:%s", global.App.Config.App.AppUrl, global.App.Config.App.Port)
	// 	Handler: app,
	// }

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// log.Println("Server exiting")

}

func handleSignal(server *gin.Engine) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGALRM)
	go func() {
		s := <-c
		log.Printf("got signal [%s],exiting now", s)
		os.Exit(0)
	}()
}
