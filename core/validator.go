package core

import (
	"errors"
	"fiber/core/response"
	"fiber/global"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var VerifyUtil = verifyUtil{}

type verifyUtil struct {
}

// VerifyForm 验证表单是否合法
func (vu verifyUtil) VerifyForm(c fiber.Ctx, obj any) error {
	bind := c.Bind()

	// 绑定查询参数
	if err := bind.Form(obj); err != nil {
		return err
	}
	// 验证查询参数
	if err := global.Validate.Struct(obj); err != nil {
		return vu.translateErr(err)
	}
	return nil
}

// VerifyJSON 验证json是否合法
func (vu verifyUtil) VerifyJSON(c fiber.Ctx, obj any) error {
	bind := c.Bind()

	// 绑定查询参数
	if err := bind.JSON(obj); err != nil {
		return err
	}
	// 验证查询参数
	if err := global.Validate.Struct(obj); err != nil {
		return vu.translateErr(err)
	}
	return nil
}

// VerifyQuery 验证查询参数是否合法
func (vu verifyUtil) VerifyQuery(ctx fiber.Ctx, obj any) error {
	bind := ctx.Bind()

	// 绑定查询参数
	if err := bind.Query(obj); err != nil {
		return response.ParamsTypeError.MakeData(err.Error())
	}
	// 验证查询参数
	err := global.Validate.Struct(obj)
	if err != nil {
		return response.ParamsValidError.MakeData(vu.translateErr(err).Error())
	}
	return nil
}

// 翻译错误信息
func (vu verifyUtil) translateErr(err error) error {
	for _, v := range err.(validator.ValidationErrors).Translate(global.Trans) {
		return errors.New(v)
	}
	return nil
}
