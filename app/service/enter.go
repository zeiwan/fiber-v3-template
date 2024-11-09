package service

import (
	"fiber/app/service/test"
	"fiber/core/uberDig"
	"fiber/gen/dao"
	"fiber/global"
	"log"
)

func InitService() {
	regFunctions := initFunctions
	regFunctions = append(regFunctions, func() *dao.Query {
		dao.SetDefault(global.Mysql)
		return dao.Q
	})
	for i := 0; i < len(regFunctions); i++ {
		if err := uberDig.ProvideForDI(regFunctions[i]); err != nil {
			log.Panic("ProvideForDI", err)
		}
	}
	log.Print("注册服务")
}

var initFunctions = []any{
	test.NewConfigService,
}
