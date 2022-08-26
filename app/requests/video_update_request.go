package requests




import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)


// ShowVideoRequest 视频详情请求
type ShowVideoRequest struct {
	ID 	  string `json:"id" form:"id" valid:"id"`
}


// ShowVideo 视频详情
func ShowVideo(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"id": []string{"required", "digits:int"},
	}

	messages := govalidator.MapData{
		"title": []string{
			"required:id 为必填项",
			"digits:id 格式错误，只能是数字且是整数",
		},
	}

	err := validate(data, rules, messages)

	return err
}
