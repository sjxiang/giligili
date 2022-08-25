package bootstrap


import (
	"os"
	"giligili/pkg/util"
)


func SetupLogger() {

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))
} 