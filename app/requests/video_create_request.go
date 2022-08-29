package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)


// CreateVideoRequest 视频投稿请求
type CreateVideoRequest struct {
	Title  string `json:"title"  form:"title"  valid:"title"`
	Info   string `json:"info"   form:"info"   valid:"info"`
	URL    string `json:"url"    form:"url"    valid:"-"`
	Avatar string `json:"avatar" form:"avatar" valid:"-"`
}


// CreateVideo 创建视频
func CreateVideo(data interface{}, c *gin.Context) map[string][]string {
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
