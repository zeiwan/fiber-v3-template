package movie

import (
	"fiber/app/schemas/req"
	"fiber/model/dao"
	"fiber/pkg/whiteCatApi"
	"fmt"
)

type IMovieService interface {
	GetHotLists(idReq req.MovieCityIdReq) (any, error)
	GetCityList() (any, error)
	GetCinemaList(listReq req.CinemaListReq) (any, error)
	GetCityArea(idReq req.MovieCityIdReq) (any, error)
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

func (s movieService) GetCityArea(idReq req.MovieCityIdReq) (any, error) {
	resp, err := s.cat.GetCityArea(idReq.CityId)
	return resp.Data, err
}

func (s movieService) GetCinemaList(listReq req.CinemaListReq) (any, error) {
	resp, err := s.cat.GetTabsCinemaList(listReq.CityId, listReq.Lng, listReq.Lat, listReq.RegionName, listReq.Keyword)
	return resp.Data, err
}

func (s movieService) GetCityList() (any, error) {
	resp, err := s.cat.CityList()
	fmt.Println(resp.Data)
	return resp.Data, err
}

func (s movieService) GetHotLists(idReq req.MovieCityIdReq) (any, error) {
	resp, err := s.cat.GetHotList(idReq.CityId)
	return resp.Data, err
}
