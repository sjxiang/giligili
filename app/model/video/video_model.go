package video

import (
	"os"

	"gorm.io/gorm"
	
	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"giligili/pkg/database"
)

// video 视频模型
type Video struct {
	gorm.Model

	Title  string
	Info   string  
	URL    string  // 视频，签名凭证中的 key
	Avatar string  // 封面，签名凭证中的 key
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


// AvatarURL 封面地址
func (videoModel *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_End_Point"), os.Getenv("OSS_AccessKey_ID"), os.Getenv("OSS_AccessKey_Secret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_Bucket"))
	signedGetURL, _ := bucket.SignURL(videoModel.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}