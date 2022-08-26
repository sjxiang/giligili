package serializer

import "giligili/app/model/video"


// VIdeo 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item video.Video) Video {
	return Video{
		ID: 		item.ID,
		Title: 		item.Title,
		Info: 		item.Info,
		CreatedAt:  item.CreatedAt.Unix(),  // Unix 时间戳
	}
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(item video.Video) Response {
	return Response{
		Msg: "创建视频成功",
		Data: BuildVideo(item),
	}
}


// BuildVideos 序列化视频列表
func BuildVideos(items []video.Video) (videos []Video) {

	for _, item := range items {

		video := BuildVideo(item)
		videos = append(videos, video)
	}

	return 
}
