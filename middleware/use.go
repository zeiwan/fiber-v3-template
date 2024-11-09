package middleware

import (
	"fiber/core/response"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// Use 初始化中间件，返回fiber.App
func Use() *fiber.App {
	conf := fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: response.ErrorHandler,
	}
	//*zap.Logger
	app := fiber.New(conf)
	// 添加日志中间件，记录请求日志
	app.Use(logger.New())
	//(fiberzap)
	//app.Use(fiberzap.New(fiberzap.NewLogger{})
	// 添加请求id中间件，生成请求id
	app.Use(requestid.New())
	// 添加跨源资源共享
	app.Use(cors.New())
	// CSRF 防御
	app.Use(csrf.New())
	// ETag
	app.Use(etag.New())
	//	添加安全头
	app.Use(helmet.New())
	//	添加幂等性中间件
	app.Use(idempotency.New())
	//	添加错误处理中间件
	app.Use(recover.New())
	//	添加自定义中间件鉴权
	app.Use(auth())

	return app
}
