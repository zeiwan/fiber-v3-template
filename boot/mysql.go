package boot

import (
	"fiber/global"
	"fiber/model/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"

	"time"
)

// InitMysql 初始化mysql数据库
func initMysql() {
	// 构建 DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		global.Conf.Mysql.Username,
		global.Conf.Mysql.Password,
		global.Conf.Mysql.Host,
		global.Conf.Mysql.Port,
		global.Conf.Mysql.Database,
		global.Conf.Mysql.Query,
	)

	// 连接到数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{Logger: gormlogger.Default})

	if err != nil {
		panic("failed to connect database")
		global.Logger.Error("初始化mysql失败", err)
		return
	}

	// 配置 GORM 日志
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 自动迁移表结构
	err = db.AutoMigrate(
		&model.GormGenTest{},
	)
	if err != nil {
		panic("failed to auto migrate")
	}

	// 设置 GORM 实例
	db.Logger = logger
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get SQL DB: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 初始化数据变更记录插件
	// _, err = loggable.Register(db, "sys_change_logs", loggable.ComputeDiff())
	if err != nil {
		global.Logger.Error("初始化mysql插件失败", err)
		return
	}
	global.Mysql = db
	global.Logger.Info("初始化mysql完成")
}
