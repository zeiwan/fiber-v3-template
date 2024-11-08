package initialize

import (
	"fiber/global"
	"fiber/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"time"
)

func initLog() {
	conf := zap.NewProductionEncoderConfig()
	// 时间格式化
	conf.EncodeTime = formatEncodeTime
	//

	filepath := filepath.Join(global.Conf.Log.Path, utils.Rename.DateName())
	logFile := &lumberjack.Logger{
		Filename:   filepath,                   // 日志文件路径
		MaxSize:    global.Conf.Log.MaxSize,    // 最大尺寸, M
		MaxBackups: global.Conf.Log.MaxBackups, // 备份数
		MaxAge:     global.Conf.Log.MaxAge,     // 存放天数
		Compress:   global.Conf.Log.Compress,   // 是否压缩
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)
	//.Sugar()
	logger := zap.New(core, zap.AddCaller())
	global.Logger = logger.Sugar()
	global.Logger.Info("初始化日志完成")
}

// formatEncodeTime 格式化 zap 日志时间
func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(time.Now().Format(time.DateTime))
}
