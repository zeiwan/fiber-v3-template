package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	AppPath string
	once    sync.Once
	GetTime times
)

type times struct {
}

func init() {
	getAppPath()
}

// 获取当前应用的路径
func getAppPath() string {
	once.Do(func() {
		// 实现获取路径的逻辑
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		AppPath = filepath.Join(dir, "admin")
	})
	return AppPath
}

// 获取13位时间戳字符串
func (t times) Get13Timestamp() string {
	// 实现获取13位时间戳的逻辑
	return strconv.FormatInt(time.Now().UnixMilli(), 10) // Go 1.17+ 支持
}
