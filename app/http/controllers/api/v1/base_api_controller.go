// 处理 v1 业务逻辑

package v1

import (
	// "giligili/app/requests"
	"giligili/pkg/oss"
	"giligili/pkg/serializer"
	// "giligili/pkg/util"

	// "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"

	// "github.com/ian-kent/go-log/log"
	// "github.com/minio/minio-go"
)


type BaseAPIController struct {
	
}


func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	})
}





// image/jpeg

func PutImage(c *gin.Context) {
	// file, _ := c.FormFile("file")
	// fileObj, err := file.Open()
	// if err != nil {
	// 	util.Log().Println(err.Error())
	// 	return	
	// }

	// ok := oss.UploadFile(file.Filename, fileObj, file.Size)
	objectName, ok := oss.UploadFile()
	if !ok {
		c.JSON(401, serializer.Response{
			Msg: "上传失败",
		})
		return
	}

	url := oss.GetFileURL(objectName)
	if url == "" {
		c.JSON(400, serializer.Response{
			Msg: "获取头像失败",
		})
		return
	}

	c.JSON(200, serializer.Response{
		Msg: "头像上传成功",
		Data: url,
	})
}


