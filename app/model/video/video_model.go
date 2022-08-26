package video

import (
	"giligili/pkg/database"

	"gorm.io/gorm"
)

// video 视频模型
type Video struct {
	gorm.Model

	Title string
	Info string
}



// Create 创建视频，通过 Video.ID 来判断是否创建成功
func (videoModel *Video) Create() {
	database.DB.Create(&videoModel)
} 

func (videoModel *Video) Show(id string) error {
	err := database.DB.First(videoModel, id).Error
	return err
}


// GetByID 通过 id 来获取视频详情
func GetByID(id string) (videoModel Video) {
	database.DB.Where("id = ?", id).First(&videoModel)
	return
}


