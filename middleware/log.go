package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

// RecordLog RouterLog 路由日志中间件
func RecordLog(ctx fiber.Ctx) error {

	switch ctx.Route().Method {
	case fiber.MethodGet:
		methodGet(ctx)
	}
	return ctx.Next()
}
func methodGet(ctx fiber.Ctx) {
	fmt.Println(ctx.Queries(), ctx.Path(), ctx.Method(), ctx.GetRespHeader("X-Request-Id"), ctx.BaseURL(), ctx.Route().Name)

}
