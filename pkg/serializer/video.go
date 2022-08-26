package serializer

import "giligili/app/model/video"


// VIdeo 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化用户
func BuildVideo(video video.Video) Video {
	return Video{
		ID: video.ID,
		Title: video.Title,
		Info: video.Info,
		CreatedAt: video.CreatedAt.Unix(),  // Unix 时间戳
	}
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(video video.Video) Response {
	return Response{
		Msg: "创建视频成功",
		Data: BuildVideo(video),
	}
}
