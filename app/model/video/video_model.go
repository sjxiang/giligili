package video

import (
	"gorm.io/gorm"
)


// video 视频模型
type Video struct {
	gorm.Model

	Title string
	Info string
}