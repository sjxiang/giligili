package serializer

import "giligili/app/model/video"


// VIdeo 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`  // Redis
	CreatedAt int64  `json:"created_at"`
}


// BuildVideo 序列化视频
func BuildVideo(item video.Video) Video {
	return Video{
		ID: 		item.ID,
		Title: 		item.Title,
		Info: 		item.Info,
		URL:        item.URL,
		Avatar:     item.AvatarURL(),  // 签名的 key 
		View:       item.View(),
		CreatedAt:  item.CreatedAt.Unix(),  // Unix 时间戳
	}
}


// BuildVideos 序列化视频列表
func BuildVideos(items []video.Video) (videos []Video) {

	for _, item := range items {

		video := BuildVideo(item)
		videos = append(videos, video)
	}

	return videos
}
