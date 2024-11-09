package middleware

import (
	"errors"
	"fiber/core/response"
	"fiber/global"
	"github.com/gofiber/fiber/v3"
)

func auth() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if err := verifyRouterAndMethod(ctx); err != nil {
			return response.Request404Error.MakeData(err.Error())
		}
		return ctx.Next()
	}
}

// verifyRouterAndMethod 验证路由和请求方法
func verifyRouterAndMethod(ctx fiber.Ctx) error {
	for _, router := range global.GetRouters {
		if router.Path == ctx.Path() && router.Method == ctx.Method() {
			return nil
		}
	}
	return errors.New("no Method or Path")
}
