package uber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"go.uber.org/dig"
)

type GroupBase struct {
	basePath    string
	initHandle  any
	regHandle   func(rg fiber.Router, group *GroupBase) error
	middlewares []fiber.Handler
}

// Group 创建了一个新的路由器组
func Group(relativePath string, initHandle any, regHandle func(rg fiber.Router, group *GroupBase) error, middlewares ...fiber.Handler) *GroupBase {
	return &GroupBase{
		basePath:    relativePath,
		initHandle:  initHandle,
		regHandle:   regHandle,
		middlewares: middlewares,
	}
}

// RegisterGroup 注册组到fiber的所有路由
func RegisterGroup(rg fiber.Router, group *GroupBase) {
	r := rg.Group(group.basePath)
	if len(group.middlewares) > 0 {
		r.Use(group.middlewares)
	}
	if err := ProvideForDI(group.initHandle); err != nil {
		log.Fatal("ProvideForDI", err)
	}
	if err := group.regHandle(r, group); err != nil {
		log.Fatal("RegHandle", err)
	}
}

// Reg 注册由DI处理的函数
func (group GroupBase) Reg(function any, opts ...dig.InvokeOption) error {
	return DI(function, opts...)
}
