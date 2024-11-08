package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
)
