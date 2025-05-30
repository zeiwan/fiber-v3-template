package movie

import (
	"fiber/app/schemas/req"
	"fiber/app/service/movie"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uber"
	"fiber/middleware"
	"github.com/gofiber/fiber/v3"
)

// MovieGroup 路由组入口（由uber框架自动加载）
var MovieGroup = uber.Group(
	"/movie",
	newMovieHandler,
	regMovieGroup,
)

// regMovieGroup 注册路由组的具体路由
func regMovieGroup(r fiber.Router, group *uber.GroupBase) error {
	return group.Reg(func(handle *movieHandler) error {
		r.Get("/hotList", handle.getHotList, middleware.RecordLog).Name("热门电影")
		r.Get("/cityList", handle.getCityList, middleware.RecordLog).Name("城市列表")
		r.Get("/getCinemaList", handle.getCinemaList, middleware.RecordLog).Name("影院列表")
		r.Get("/getCityArea", handle.getCityArea, middleware.RecordLog).Name("城市区域")
		return nil
	})
}

// newMovieHandler 构造函数
func newMovieHandler(srv movie.IMovieService) *movieHandler {
	return &movieHandler{srv: srv}
}

// movieHandler 路由处理器
type movieHandler struct {
	srv movie.IMovieService
}

func (h movieHandler) getCityList(ctx fiber.Ctx) error {
	resp, err := h.srv.GetCityList()
	return response.CheckAndRespWithData(ctx, resp, err)
}

func (h movieHandler) getHotList(ctx fiber.Ctx) error {
	var obj req.MovieCityIdReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}
	resp, err := h.srv.GetHotLists(obj)
	return response.CheckAndRespWithData(ctx, resp, err)
}

func (h movieHandler) getCinemaList(ctx fiber.Ctx) error {
	var obj req.CinemaListReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}
	resp, err := h.srv.GetCinemaList(obj)
	return response.CheckAndRespWithData(ctx, resp, err)
}

func (h movieHandler) getCityArea(ctx fiber.Ctx) error {
	var obj req.MovieCityIdReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}
	resp, err := h.srv.GetCityArea(obj)

	return response.CheckAndRespWithData(ctx, resp, err)
}
