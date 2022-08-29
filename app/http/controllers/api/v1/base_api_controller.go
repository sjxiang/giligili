// 处理 v1 业务逻辑

package v1

import (
	"os"

	"giligili/pkg/serializer"


	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type BaseAPIController struct {
	
}


func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	})
}


func UploadToken(c *gin.Context) {
	c.JSON(200, Post(c))
}


func Post(c *gin.Context) serializer.Response {

	client, err := oss.New(os.Getenv("OSS_End_Point"), os.Getenv("OSS_AccessKey_ID"), os.Getenv("OSS_AccessKey_Secret"))
	if err != nil {		
		return serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		}
	}

	// 获取存储空间
	bucket, err := client.Bucket(os.Getenv("OSS_Bucket"))
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		} 
	}

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType("image/jpeg"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ".jpg"

	// 签名直传
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)  
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		} 
	}

	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code: 50002,
			Msg: "OSS 配置错误",
			Error: err,
		} 
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,  
			"get": signedGetURL,  
		},
	}
}



