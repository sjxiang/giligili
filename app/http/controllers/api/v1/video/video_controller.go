package video

import (
	"net/http"

	"giligili/app/http/controllers/api/v1"
	"giligili/app/model/video"
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
	videoModel := video.Video{
		Title: request.Title,
		Info: request.Info,
	}
	videoModel.Create()

	if videoModel.ID > 0 {
		c.JSON(http.StatusOK, serializer.BuildVideoResponse(videoModel))
		return
	}

	c.JSON(http.StatusInternalServerError, serializer.Response{  // 500 服务器内部错误
		Code: 50001,
		Msg: "视频保存失败，请稍后再试。",
	})
}


// ShowVideo 视频详情（参数校验，TODO .. ）
func (vc VideoController) ShowVideo(c *gin.Context) {

	id := c.Param("id")
	var videoModel video.Video
	
	err := videoModel.Show(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 404,
			Msg: "视频不存在", 
			Error: err,
		})
		
		return 
	}

	c.JSON(http.StatusOK, serializer.BuildVideoResponse(videoModel))

	// videoModel := video.GetByID(id)

	// if videoModel.ID == 0 {
	// 	  c.JSON(http.StatusInternalServerError, serializer.ParamErr("视频不存在", nil,))
	//    return 
	// }
}


// ListVideo 视频列表
func (vc VideoController) ListVideo(c *gin.Context) {



}



// UpdateVideo 视频更新
func (vc VideoController) UpdateVideo(c *gin.Context) {



}


// DeleteVideo 视频删除   
func (vc VideoController) DeleteVideo(c *gin.Context) {



}
