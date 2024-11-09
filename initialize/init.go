package initialize

func init() {
	initConfig()

	initLog()
	initValidate()

	initMysql()
	initRedis()
}
