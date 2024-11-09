package utils

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	AppPath string
	once    sync.Once
)

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
		AppPath = filepath.Join(dir, "app")
	})
	return AppPath
}
