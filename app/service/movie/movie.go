package movie

import (
	"fiber/app/schemas/req"
	"fiber/model/dao"
	"fiber/pkg/whiteCatApi"
)

type IMovieService interface {
	GetMovieHotLists(listReq req.MovieHotListReq) (any, error)
}

func NewMovieService() IMovieService {
	return &movieService{
		cat: whiteCatApi.NewCatClient(),
	}
}

type movieService struct {
	dao *dao.Query
	cat *whiteCatApi.CatClient
}

func (s movieService) GetMovieHotLists(listReq req.MovieHotListReq) (any, error) {
	//var obj []resp.MovieHotListResp
	s.cat.GetCityList()
	return nil, nil
}
