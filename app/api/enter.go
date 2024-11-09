package api

import (
	"fiber/app/api/test"
	"fiber/core/uberDig"
	"fiber/global"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log"
)

func InitRouter(app *fiber.App) {
	// 指定路由前缀
	group := app.Group(global.Conf.Server.URLPrefix)
	routers := InitRouters[:]
	fmt.Println("注册路由开始", routers)
	for i := 0; i < len(routers); i++ {
		fmt.Println("注册路由", i, routers[i])
		uberDig.RegisterGroup(group, routers[i])
	}
	log.Print("注册路由完成")
}

var InitRouters = []*uberDig.GroupBase{
	test.ConfigGroup,
}
