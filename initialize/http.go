package initialize

import (
	"fiber/app/api"
	"fiber/app/service"
	"fiber/global"
	"fiber/middleware"
	"fmt"
)

func InitHttpServer() {
	app := middleware.Use()
	// 初始化服务
	service.InitService()

	// 初始化路由
	api.InitRouter(app)

	// 打印所有路由
	Listen := fmt.Sprintf(":%d", global.Conf.Port)
	app.Listen(Listen)
	global.Logger.Fatal()
}
