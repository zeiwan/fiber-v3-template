package response

import (
	"fiber/global"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// RespType 响应类型
type RespType struct {
	code int
	msg  string
	data interface{}
	show int // 0:不显示 1:显示
}

// Response 响应格式结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Show int         `json:"show"`
}

var (
	Success  = RespType{code: 1, msg: "成功", show: 0}
	Failed   = RespType{code: 0, msg: "失败", show: 1}
	FileDown = RespType{code: 2, msg: "文件下载", show: 1}

	ParamsValidError    = RespType{code: 0, msg: "参数校验错误"}
	ParamsTypeError     = RespType{code: 0, msg: "参数类型错误"}
	RequestMethodError  = RespType{code: 0, msg: "请求方法错误"}
	AssertArgumentError = RespType{code: 0, msg: "断言参数错误"}

	QueryError  = RespType{code: 0, msg: "查询资源失败"}
	DeleteError = RespType{code: 0, msg: "删除资源失败"}
	EditError   = RespType{code: 0, msg: "编辑数据失败"}
	CreateError = RespType{code: 0, msg: "创建数据失败"}

	LoginAccountError = RespType{code: 0, msg: "登录账号或密码错误"}
	LoginDisableError = RespType{code: 0, msg: "登录账号已被禁用了"}
	TokenEmpty        = RespType{code: 0, msg: "token参数为空"}
	TokenInvalid      = RespType{code: 0, msg: "token参数无效"}

	NoPermission    = RespType{code: 0, msg: "无相关权限"}
	Request404Error = RespType{code: 0, msg: "请求接口不存在"}
	Request405Error = RespType{code: 0, msg: "请求方法不允许"}

	SystemError = RespType{code: 0, msg: "系统错误"}
)

// Error 实现error方法
func (rt RespType) Error() string {
	return strconv.Itoa(rt.code) + ":" + rt.msg
}

// Make 以响应类型生成信息
func (rt RespType) Make(msg string) RespType {
	rt.msg = msg
	return rt
}

// MakeData 以响应类型生成数据
func (rt RespType) MakeData(data interface{}) RespType {
	rt.data = data
	return rt
}

// Code 获取code
func (rt RespType) Code() int {
	return rt.code
}

// Msg 获取msg
func (rt RespType) Msg() string {
	return rt.msg
}

// Data 获取data
func (rt RespType) Data() interface{} {
	return rt.data
}

// Result 统一响应
func Result(resp RespType) Response {
	fmt.Println(resp.msg, "====RespType======")
	//show := 0
	//if !errors.Is(resp, Success) {
	//	show = 1
	//}
	//
	//if resp != Success {
	//	fmt.Println(resp.msg, "====err")
	//}

	return Response{
		Code: resp.code,
		Msg:  resp.msg,
		Data: resp.data,
		Show: 1,
	}
}

// CheckAndResp 判断是否出现错误，并返回对应响应
func CheckAndResp(ctx fiber.Ctx, err error) error {
	_ = ErrorHandler(ctx, err)
	return ctx.JSON(Result(Success))
}

// CheckAndRespWithData 判断是否出现错误，并返回对应响应（带data数据）
func CheckAndRespWithData(ctx fiber.Ctx, data interface{}, err error) error {
	err = ErrorHandler(ctx, err)
	fmt.Println(err, "====CheckAndRespWithData")
	if err != nil {
		fmt.Println(err, "====CheckAndRespWithData", err)

		return err
	}

	return ctx.JSON(Result(Success.MakeData(data)))

}

// // 下载
//
//	func CheckDownloadResp(c *gin.Context, data any, err error) {
//		if IsFailWithResp(c, err) {
//			return
//		}
//		Result(c, FileDown, data)
//	}
//
// FailLog 错误响应
func FailLog(c fiber.Ctx, resp RespType) error {
	loggerResp(resp, "Request Fail: url=[%s], resp.tpl=[%+v]", c.Path(), resp)
	return c.JSON(Result(resp))
}

// FailWithDataLog 错误响应附带data
func FailWithDataLog(ctx fiber.Ctx, resp RespType) error {
	loggerResp(resp, "Request FailWithData: url=[%s], resp.tpl=[%+v], data=[%+v]", ctx.Path(), resp.msg, resp.data)
	return ctx.JSON(Result(resp))
}

// loggerResp 把错误写入日志
func loggerResp(resp RespType, template string, args ...interface{}) {
	loggerFunc := global.Logger.Warnf
	if resp.code >= 500 {
		loggerFunc = global.Logger.Errorf
	}
	loggerFunc(template, args...)
}

// ErrorHandler 统一处理错误
func ErrorHandler(ctx fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}
	var v RespType
	fmt.Println(err.(RespType), "====ErrorHandler")
	data := v.Data()
	if data == nil {
		data = []string{}
	}
	return ctx.JSON(fiber.Map{"code": v.Code(), "msg": v.Msg(), "data1": data})
}