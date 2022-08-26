package video

import (
	"giligili/app/http/controllers/api/v1"
	"giligili/app/requests"
	"giligili/pkg/serializer"

	"github.com/gin-gonic/gin"
)

// 视频控制器
type VideoController struct {
	v1.BaseAPIController
}


// CreateVideo 视频投稿
func (v VideoController) CreateVideo(c *gin.Context) {
	
	// 1. 验证表单
	request := requests.CreateVideoRequest{}
	if ok := requests.Validate(c, &request, requests.CreateVideo); !ok {
		return
	}

	// 2. 验证成功，创建数据

	
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg: "成功",
	})
}