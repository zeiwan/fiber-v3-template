package login

import (
	"fiber/admin/schemas/req"
	"fiber/model/dao"
	"fmt"
)

type ILoginService interface {
	Login(loginReq req.LoginReq) (any, error)
}

func NewLoginService(dao *dao.Query) ILoginService {
	return &loginService{
		dao: dao,
	}
}

type loginService struct {
	dao *dao.Query
}

func (s loginService) Login(loginReq req.LoginReq) (any, error) {
	m := s.dao.GormGenTest
	fmt.Println(m)
	//q := m.WithContext(context.Background())
	//first, err := q.Where(m.Account.Eq(loginReq.Account)).First()
	return nil, nil
}
