// 处理 v1 业务逻辑

package v1

import (
	"giligili/app/requests"
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



// Post 创建 Token
func UploadToken(c *gin.Context) {
	
	// 1. 验证表单
	request := requests.UploadTokenRequest{}
	if ok := requests.Validate(c, &request, requests.UploadToken); !ok {
		return
	}

	// 2. 验证成功，发送 Token
	requests.SendToken(c, request)
	
}

