package movie

import (
	"fiber/app/schemas/req"
	"fiber/model/dao"
	"fiber/pkg/whiteCatApi"
	"fmt"
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
	err, resp := s.cat.GetHotList(listReq.CityId)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Data)
	return resp, nil
}
