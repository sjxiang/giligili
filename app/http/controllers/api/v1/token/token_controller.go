package token

import (
	"net/http"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"giligili/app/http/controllers/api/v1"
	"giligili/pkg/serializer"
)

// Token 控制器
type TokenController struct {
	v1.BaseAPIController
}


func (tc TokenController) UploadToken(c *gin.Context) {

	// 初始化 oss 客户端
	client, err := oss.New(os.Getenv("OSS_End_Point"), os.Getenv("OSS_AccessKey_ID"), os.Getenv("OSS_AccessKey_Secret"))
	if err != nil {		

		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		})
		return 
	}

	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_Bucket"))
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
		oss.ContentType("image/jpeg"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ".jpg"

	// 签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)  // 600 秒
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

	c.JSON(http.StatusOK, serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,  // 上传图片凭证
			"get": signedGetURL,  // 未来下载图片凭证
		},
	})
}



