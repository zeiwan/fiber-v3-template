package test

import (
	"fiber/app/schemas/req"
	"fiber/app/service/test"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uber"
	"fiber/middleware"
	"github.com/gofiber/fiber/v3"
)

// ConfigGroup 路由组入口（由uber框架自动加载）
var ConfigGroup = uber.Group(
	"/test",
	newConfigHandler,
	regConfigGroup)

// regConfigGroup 注册路由组的具体路由
func regConfigGroup(r fiber.Router, group *uber.GroupBase) error {
	return group.Reg(func(handle *configHandler) error {
		// 示例路由：GET /test/get
		r.Get("/get", handle.getConfig, middleware.RecordLog).Name("获取配置")
		return nil
	})
}

// newConfigHandler 依赖注入构造函数
func newConfigHandler(srv test.IConfigService) *configHandler {
	return &configHandler{srv: srv}
}

// configHandler 路由处理器
type configHandler struct {
	srv test.IConfigService
}

// getConfig 具体处理逻辑
func (h configHandler) getConfig(ctx fiber.Ctx) error {
	// 1. 请求参数校验
	var obj req.TestGetReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err // 自动返回400错误（由core.VerifyUtil处理）
	}

	// 2. 调用Service层
	resp, err := h.srv.GetConfig(obj)

	// 3. 统一响应格式（成功/错误）
	return response.CheckAndRespWithData(ctx, resp, err)
}
