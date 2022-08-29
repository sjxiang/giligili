package main

import (
	"giligili/bootstrap"
	"giligili/conf"
	"giligili/pkg/util"

	"github.com/gin-gonic/gin"
)


func init() {

	// 从配置文件读取配置
	conf.Init()
	
	// 初始化日志
	bootstrap.SetupLogger()
	
	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化 mysql
	bootstrap.SetUpDB()

}


func main() {
	
	router := gin.New()

	// 装载路由
	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		util.Log().Panic(err.Error())
	}

}
