package core

import (
	"errors"
	"fiber/core/response"
	"fiber/global"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var VerifyUtil = verifyUtil{}

type verifyUtil struct{}

// bindAndValidate 绑定并验证数据
func (vu verifyUtil) bindAndValidate(obj any, bindFunc func(any) error) error {
	// 绑定数据
	if err := bindFunc(obj); err != nil {
		return response.ParamsTypeError.MakeData(vu.translateErr(err).Error())
	}
	// 验证数据
	if err := global.Validate.Struct(obj); err != nil {
		return response.ParamsValidError.MakeData(vu.translateErr(err).Error())
	}
	return nil
}

// VerifyForm 验证表单是否合法
func (vu verifyUtil) VerifyForm(ctx fiber.Ctx, obj any) error {
	return vu.bindAndValidate(obj, ctx.Bind().Form)
}

// VerifyJSON 验证json是否合法
func (vu verifyUtil) VerifyJSON(ctx fiber.Ctx, obj any) error {
	return vu.bindAndValidate(obj, ctx.Bind().JSON)
}

// VerifyQuery 验证查询参数是否合法
func (vu verifyUtil) VerifyQuery(ctx fiber.Ctx, obj any) error {
	return vu.bindAndValidate(obj, ctx.Bind().Query)
}

// translateErr 翻译错误信息
func (vu verifyUtil) translateErr(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErrors.Translate(global.Trans) {
			return errors.New(v)
		}
	}
	return err
}
