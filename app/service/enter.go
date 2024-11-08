package service

import (
	"fiber/app/service/test"
	"fiber/core/uberDig"
	"log"
)

func InitService() {
	regFunctions := initFunctions
	for i := 0; i < len(regFunctions); i++ {
		if err := uberDig.ProvideForDI(regFunctions[i]); err != nil {
			log.Panic("ProvideForDI", err)
		}
	}
	log.Print("注册服务")
}

var initFunctions = []interface{}{
	test.NewConfigService,
}
