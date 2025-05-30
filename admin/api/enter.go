package api

import (
	"fiber/admin/api/login"
	"fiber/core/uber"
	"fiber/global"
	"github.com/gofiber/fiber/v3"
	"log"
)

func InitRouter(app *fiber.App) {
	// 指定路由前缀
	group := app.Group(global.Conf.Server.AdminPrefix)
	routers := initRouters[:]
	for i := 0; i < len(routers); i++ {
		uber.RegisterGroup(group, routers[i])
	}
	log.Print("注册路由完成")
}

var initRouters = []*uber.GroupBase{
	login.LoginGroup,
}
