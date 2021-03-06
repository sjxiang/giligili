package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {  // 绑定 binding
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// ShowVideo 视频详情
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err == nil {  // 绑定 binding
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// ListVideo 视频列表
func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {  // 绑定 binding
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}



// UpdateVideo 更新视频详情
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {  // 绑定 binding
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}



// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	if err := c.ShouldBind(&service); err == nil {  // 绑定 binding
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
