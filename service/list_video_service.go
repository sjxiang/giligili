package service

import (
	"giligili/model"
	"giligili/serializer"
)

//  ListVideoService 视频列表的服务
type ListVideoService struct {
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	
	if err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:  "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}