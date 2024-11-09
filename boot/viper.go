package boot

import (
	"fiber/global"
	"fmt"
	"github.com/gofiber/fiber/v3/log"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置文件函数
func initConfig() {
	// 设置配置文件名称（不包括扩展名）
	viper.SetConfigName("config")
	// 设置配置文件格式为yaml
	viper.SetConfigType("yaml")
	// 添加当前目录为配置查找路径
	viper.AddConfigPath("conf")
	//	读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	//    将配置信息绑定到结构体体中
	err = viper.Unmarshal(&global.Conf)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	fmt.Println(global.Conf)
}
