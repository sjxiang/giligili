package routes

import (
	"giligili/app/http/controllers/api/v1/video"
	base "giligili/app/http/controllers/api/v1"

	"github.com/gin-gonic/gin"
)


func RegisterApiRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", base.Ping)


		vc := new(video.VideoController)

		v1.POST("/videos", vc.CreateVideo)
		v1.GET("/video/:id", vc.ShowVideo)
		v1.GET("/videos", vc.ListVideo)
		v1.PUT("/video/:id", vc.UpdateVideo)
		v1.DELETE("/video/:id", vc.DeleteVideo)
	
	}
}

	// // 中间件, 顺序不能改
	// r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// r.Use(middleware.Cors())
	// r.Use(middleware.CurrentUser())


	// POST localhost:3000/api/v1/ping 存活检查

	// 路由
	// v1 := r.Group("/api/v1")
	// {
	// 	v1.POST("ping", api.Ping)

	// 	// 用户登录
	// 	v1.POST("user/register", api.UserRegister)

	// 	// 用户登录
	// 	v1.POST("user/login", api.UserLogin)

	// 	// 需要登录保护的
	// 	auth := v1.Group("")
	// 	auth.Use(middleware.AuthRequired())
	// 	{
	// 		// User Routing
	// 		auth.GET("user/me", api.UserMe)
	// 		auth.DELETE("user/logout", api.UserLogout)
	// 	}
	// }
	
	
// }
