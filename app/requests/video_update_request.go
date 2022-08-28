package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)


// UpdateVideoRequest 视频更新请求
type UpdateVideoRequest struct {
	Title string `json:"title" form:"title" valid:"title"`
	Info  string `json:"info"  form:"info"  valid:"info"`
}


// UpdateVideo 视频更新
func UpdateVideo(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min:2", "max:30"},
		"info":  []string{ "min:0", "max:200"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required: 标题为必填项",
			"min: 长度需大于 2",
			"max: 长度需小于 30",
		},
		"info": []string{
			"min: 可以不填，为空",
			"max: 长度需小于 200",
		},
	}

	err := validate(data, rules, messages)

	return err
}
