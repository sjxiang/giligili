package requests

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"

	"giligili/pkg/serializer"
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





func SendToken(c *gin.Context, utr UploadTokenRequest) {

	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {		
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		})

		return 
	}

	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		})

		return 
	}

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ".png"

	// 签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)  
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		})

		return 
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		})

		return 
	}

	c.JSON(200, serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,  // 上传地址
			"get": signedGetURL,  // 下载地址
		},
	})
}

