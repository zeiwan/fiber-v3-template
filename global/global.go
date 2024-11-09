package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 优雅的初始化配置
var (
	// Conf  系统配置
	Conf config
	// Logger  日志
	Logger *zap.SugaredLogger
	//	Trans  验证器
	Trans    ut.Translator
	Validate *validator.Validate
	// Mysql  数据库
	Mysql *gorm.DB
	// Redis  缓存
	Redis *redis.Client
	// GetRouters	路由
	GetRouters []fiber.Route
)
