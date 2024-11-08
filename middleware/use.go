package middleware

import (
	"fiber/core/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	jsoniter "github.com/json-iterator/go"
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
		JSONDecoder:  jsoniter.Unmarshal,
		JSONEncoder:  jsoniter.Marshal,
		ErrorHandler: response.ErrorHandler,
		//DisableKeepalive: true,
		//StructValidator: &structValidator{validate: validator.New()},
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

	//app.Use =  // 注意：这里没有括号，因为是传递函数本身
	return app
}

// Response 定义了错误响应的结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
