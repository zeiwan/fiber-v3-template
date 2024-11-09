package test

import (
	"fiber/app/schemas/req"
	"fiber/app/schemas/resp"
	"fiber/core/response"
)

type IConfigService interface {
	GetConfig(req.TestGetReq) (any, error)
}

func NewConfigService() IConfigService {
	return &configService{}
}

type configService struct {
	//dao *dao.Query
}

func (c configService) GetConfig(tgReq req.TestGetReq) (any, error) {
	var obj resp.TestGetResp
	response.Copy(&obj, tgReq)
	return obj, nil
}
