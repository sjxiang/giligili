package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

)

// UploadToken 获得上传 oss token 的请求
type UploadTokenRequest struct {
	Filename string `json:"filename" form:"filename" valid:"filename"`
}


func UploadToken(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"filename": []string{"required"},
	}

	messages := govalidator.MapData{
		"filename": []string{
			"required: 标题为必填项",
		},
	}

	err := validate(data, rules, messages)

	return err
}



