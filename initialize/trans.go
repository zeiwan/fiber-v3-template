package initialize

import (
	"fiber/global"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func initValidate() {
	// 初始化验证器
	validate = validator.New()

	// 初始化翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	// 注册中文翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		global.Logger.Fatalf("初始化验证器翻译器失败: %v", err)
		return
	}
	global.Validate = validate
	global.Trans = trans
	global.Logger.Info("初始化验证器翻译器成功")
}
