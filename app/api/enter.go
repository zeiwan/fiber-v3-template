package api

import (
	"fiber/app/api/movie"
	"fiber/app/api/test"
	"fiber/core/uber"
	"fiber/global"
	"github.com/gofiber/fiber/v3"
	"log"
)

func InitRouter(app *fiber.App) {
	// 指定路由前缀
	group := app.Group(global.Conf.Server.URLPrefix)
	routers := InitRouters[:]
	for i := 0; i < len(routers); i++ {
		uber.RegisterGroup(group, routers[i])
	}
	log.Print("注册路由完成")
}

var InitRouters = []*uber.GroupBase{
	test.ConfigGroup,
	movie.MovieGroup,
}
