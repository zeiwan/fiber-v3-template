package movie

import (
	"fiber/app/schemas/req"
	"fiber/app/service/movie"
	"fiber/core"
	"fiber/core/response"
	"fiber/core/uber"
	"fiber/middleware"
	"fmt"
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
		// 路由：GET /movie/get
		r.Get("/get", handle.getMovieHotList, middleware.RecordLog).Name("获取电影热门列表")
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

// getMovie 处理逻辑
func (h movieHandler) getMovieHotList(ctx fiber.Ctx) error {
	var obj req.MovieHotListReq
	if err := core.VerifyUtil.VerifyQuery(ctx, &obj); err != nil {
		return err
	}
	fmt.Println(obj)
	//resp, err := h.srv.GetMovieHotLists(obj)
	return response.CheckAndRespWithData(ctx, nil, nil)
}
