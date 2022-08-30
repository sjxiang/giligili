package video

import (
	"net/http"

	"giligili/app/http/controllers/api/v1"
	"giligili/app/model/video"
	"giligili/app/requests"
	"giligili/pkg/database"
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
		Title:  request.Title,
		Info:   request.Info,
		URL:    request.URL,
		Avatar: request.Avatar,
	}
	videoModel.Create()

	if videoModel.ID > 0 {
		c.JSON(http.StatusOK, serializer.Response{
			Msg: "创建视频成功",
			Data: serializer.BuildVideo(videoModel),
		})
		
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
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: 404,
			Msg: "视频不存在", 
			Error: err,
		})
		
		return 
	}

	c.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildVideo(videoModel),
	})
	
	// videoModel := video.GetByID(id)

	// if videoModel.ID == 0 {
	// 	  c.JSON(http.StatusInternalServerError, serializer.ParamErr("视频不存在", nil,))
	//    return 
	// }
}


// ListVideo 视频列表
func (vc VideoController) ListVideo(c *gin.Context) {

	var videoModels []video.Video
	
	err := database.DB.Find(&videoModels).Error  // （应该区分错误。数据库查询错误、语法错误 ... ）
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50000,
			Msg: "数据库连接错误",
			Error: err,
		})

		return 
	}
 
	c.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildVideos(videoModels),
	})
}



// UpdateVideo 视频更新
func (vc VideoController) UpdateVideo(c *gin.Context) {

	// 1. 先查询，有没有这个 video 投稿	
	id := c.Param("id")
	var videoModel video.Video
	
	err := videoModel.Show(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: 404,
			Msg: "视频不存在", 
			Error: err,
		})
		
		return 
	}

	
	// 2. 验证表单
	request := requests.UpdateVideoRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateVideo); !ok {
		return
	}

	// 3. 验证成功，更新数据
	videoModel.Title = request.Title
	videoModel.Info = request.Info
	database.DB.Save(&videoModel)

	err = database.DB.Save(&videoModel).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{  // 500 服务器内部错误
			Code: 50001,
			Msg: "视频保存失败，请稍后再试。",
		})

		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildVideo(videoModel),
	})
}


// DeleteVideo 视频删除   
func (vc VideoController) DeleteVideo(c *gin.Context) {
	
	// 1. 先查询，有没有这个 video 投稿	
	id := c.Param("id")
	var videoModel video.Video

	err := videoModel.Show(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code: 404,
			Msg: "视频不存在", 
			Error: err,
		})
		
		return 
	}

	// 2. 删除
	err = database.DB.Delete(&videoModel).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{  // 500 服务器内部错误
			Code: 50001,
			Msg: "视频保存失败，请稍后再试。",
		})

		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Msg: "删除视频成功",
	})
}
