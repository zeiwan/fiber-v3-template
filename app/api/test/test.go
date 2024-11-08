package test

import (
	"fiber/app/service/test"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uberDig"
	"fmt"
	"github.com/gofiber/fiber/v3"
)

var ConfigGroup = uberDig.Group("/config", newConfigHandler, regConfigGroup)

func regConfigGroup(r fiber.Router, group *uberDig.GroupBase) error {
	// 在这里添加实际的路由注册逻辑
	return group.Reg(func(handle *configHandler) error {
		r.Get("/getConfig", handle.getConfig, log)
		return nil
	})
}

func log(ctx fiber.Ctx) error {
	fmt.Println("path", ctx.Path())
	//ctx.Path()
	err := ctx.Next()
	if err != nil {
		return err
	}
	return nil
}
func newConfigHandler(srv test.IConfigService) *configHandler {
	return &configHandler{srv: srv}
}

type configHandler struct {
	srv test.IConfigService
}

func (h configHandler) getConfig(ctx fiber.Ctx) error {
	var obj Req

	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}

	_, err := h.srv.GetConfig()
	fmt.Println("obj", err)
	return response.CheckAndResp(ctx, err)
}

type Req struct {
	Account string `form:"account"`
	Id      int    `form:"id" validate:"required"`
	Ids     int    `form:"ids" validate:"required"`
}
