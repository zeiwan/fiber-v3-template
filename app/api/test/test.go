package test

import (
	"fiber/app/schemas/req"
	"fiber/app/service/test"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uberDig"
	"fiber/middleware"
	"github.com/gofiber/fiber/v3"
)

var ConfigGroup = uberDig.Group("/test", newConfigHandler, regConfigGroup)

func regConfigGroup(r fiber.Router, group *uberDig.GroupBase) error {
	// 在这里添加实际的路由注册逻辑
	return group.Reg(func(handle *configHandler) error {
		r.Get("/get", handle.getConfig, middleware.RecordLog).Name("优雅的路由")
		return nil
	})
}

func newConfigHandler(srv test.IConfigService) *configHandler {
	return &configHandler{srv: srv}
}

type configHandler struct {
	srv test.IConfigService
}

func (h configHandler) getConfig(ctx fiber.Ctx) error {
	var obj req.TestGetReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}
	resp, err := h.srv.GetConfig(obj)
	return response.CheckAndRespWithData(ctx, resp, err)
}
