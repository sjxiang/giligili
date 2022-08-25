package video

import (
	"giligili/app/http/controllers/api/v1"
	"giligili/pkg/serializer"

	"github.com/gin-gonic/gin"
)

// 视频控制器
type VideoController struct {
	v1.BaseAPIController
}


// CreateVideo 视频投稿
func (v VideoController) CreateVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg: "成功",
	})
}