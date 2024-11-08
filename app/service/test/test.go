package test

import "fiber/core/response"

type IConfigService interface {
	GetConfig() (string, error)
}

func NewConfigService() IConfigService {
	return &configService{}
}

type configService struct {
	//dao *dao.Query
}

func (c configService) GetConfig() (string, error) {
	return "hello world", response.QueryError
}
