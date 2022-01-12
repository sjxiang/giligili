package server

import (
	"os"
	"giligili/api"
	"giligili/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())  // 跨域
	r.Use(middleware.CurrentUser()) // 验证登录

	// 路由
	v1 := r.Group("/api/v1")
	{

		v1.POST("videos", api.CreateVideo)     // 投稿视频  
		// curl "http://localhost:3000/api/v1/videos" 
		// -H "Content-Type: application/json" 
		// -d "{\"status\": 0, \"data\": null, \"msg\": \"pong\", \"error\": \"\"}" 
		// -X POST
		
		// v1.GET("video/:id", api.ShowVideo)  // 视频详情
		// v1.GET("videos", api.ListVideo)    // 视频列表
		// v1.PUT("video/id", api.UpdateVideo) // 更新视频详情
		// v1.DELETE("video/:id", api.DeleteVideo) // 删除视频


		v1.POST("ping", api.Ping)    // localhost:3000/api/v1/ping 存活检查

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

	}
	return r
}
