package boot

func init() {
	initConfig()

	initLog()
	initValidate()

	initMysql()
	initRedis()
}
