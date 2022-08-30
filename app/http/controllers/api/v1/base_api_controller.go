// 处理 v1 业务逻辑

package v1

import (
	"giligili/pkg/serializer"


	"github.com/gin-gonic/gin"
)


type BaseAPIController struct {
	
}


func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	})
}


