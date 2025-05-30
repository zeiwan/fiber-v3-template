package test

import (
	"context"
	"fiber/app/schemas/req"
	"fiber/app/schemas/resp"
	"fiber/core/response"
	"fiber/model/dao"
	"github.com/gofiber/fiber/v3"
)

type IConfigService interface {
	GetConfig(fiber.Ctx, req.TestGetReq) (any, error)
}

func NewConfigService(dao *dao.Query) IConfigService {
	return &configService{
		dao: dao,
	}
}

type configService struct {
	dao *dao.Query
}

func (c configService) GetConfig(ctx fiber.Ctx, tgReq req.TestGetReq) (any, error) {
	var obj resp.TestGetResp
	response.Copy(&obj, tgReq)
	m := c.dao.GormGenTest
	q := m.WithContext(context.Background())
	first, err := q.Where(m.ID.Eq(int32(tgReq.Id))).First()
	if err != nil {
		return nil, err
	}
	panic(first)
	return obj, nil
}
