package login

import (
	"fiber/admin/schemas/req"
	"fiber/admin/service/login"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uber"
	"fiber/middleware"
	"github.com/gofiber/fiber/v3"
)

// LoginGroup 路由组入口（由uber框架自动加载）
var LoginGroup = uber.Group(
	"/login",
	newLoginHandler,
	regLoginGroup,
)

// regLoginGroup 注册路由组的具体路由
func regLoginGroup(r fiber.Router, group *uber.GroupBase) error {
	return group.Reg(func(handle *loginHandler) error {
		r.Post("/account", handle.login, middleware.RecordLog).Name("Login")
		return nil
	})
}

// newLoginHandler 构造函数
func newLoginHandler(srv login.ILoginService) *loginHandler {
	return &loginHandler{srv: srv}
}

// loginHandler 路由处理器
type loginHandler struct {
	srv login.ILoginService
}

// login 处理逻辑
func (h loginHandler) login(ctx fiber.Ctx) error {
	var obj req.LoginReq
	if err := core.VerifyUtil.VerifyJSON(ctx, &obj); err != nil {
		return err
	}
	resp, err := h.srv.Login(obj)
	return response.CheckAndRespWithData(ctx, resp, err)
}
