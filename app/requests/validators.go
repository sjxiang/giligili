package requests

import (
	"giligili/pkg/serializer"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string


func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {


	// 1. 解析请求，支持 JSON 数据
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(200, serializer.Err(
			40000,
			"请求解析错误，请确认格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			err,
		))	

		return false
	}


	// 2，表单验证
	errs := handler(obj, c)	// 回调、钩子

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		c.JSON(200, serializer.Response{
			Code: serializer.CodeParamErr,
			Msg: "请求验证不通过，具体请查看 errors",
			Error: errs,
		})

		return false
	}

	return true
}


func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置选项
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
