package oss

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"giligili/pkg/util"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/policy"
	"github.com/google/uuid"
)

var MinioClient *minio.Client

func InitMinio() error {

	// 步骤一：创建 minio 客户端对象
	// minio.New 参数
	// 1. 对象存储服务的 URL
	// 2. 略
	// 3. 略
	// 4. 是否使用 HTTPS

	var err error
	
	MinioClient, err := minio.New(
		os.Getenv("MINIO_END_POINT"),  // 对象存储服务的 URL
		os.Getenv("MINIO_ROOT_USER"), 
		os.Getenv("MINIO_ROOT_PASSWORD"), 
		false,  // 是否使用 HTTPS
	)
	if err != nil {	
		return err
	}


	// 步骤二：创建存储桶 
	err = MinioClient.MakeBucket(os.Getenv("MINIO_BUCKET"), "cn-ShangHai-01")
	if err != nil {
		
		// 检查存储桶是否已经存在
		exists, err := MinioClient.BucketExists(os.Getenv("MINIO_BUCKET"))
		if err != nil || !exists {
			return err
		}
	}

	
	// 步骤三：设置策略
	err = MinioClient.SetBucketPolicy(os.Getenv("MINIO_BUCKET"), policy.BucketPolicyReadWrite)
	if err != nil {
		return err
	}

	return nil 
}


// 	// 步骤四：上传文件
// 	n, err := client.PutObject(os.Getenv("MINIO_BUCKET"),)
	
// 	// 上传文件名
//     objectName := uuid.Must(uuid.NewRandom()).String() + ".jpg"
// 	// 上传文件路径
//     // filePath := "/upload/avatar/" + objectName
//     // 上传文件类型
// 	// contentType := "application/image"

//     // // 使用 FPutObject 上传一个 jpg 文件。
//     // n, err := client.FPutObject(os.Getenv("OSS_BUCKET"), objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
//     // if err != nil {
//     //     log.Fatalln(err)
//     // }
	
// 	// log.Printf("步骤 3. 上传文件 %s 成功\n", objectName)


// 	// 签名直传
// 	expiry := time.Second * 24 * 60 * 60 // 1 day.
// 	signedPutURL, err := client.PresignedPutObject(os.Getenv("OSS_BUCKET"), objectName, expiry)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, serializer.Response{
// 			Code: 50002,
// 			Msg: "OSS 配置错误，上传 URL",
// 			Error: err,
// 		})

// 		return 
// 	}

// 	reqParams := make(url.Values)

// 	signedGetURL, err := client.PresignedGetObject(os.Getenv("OSS_BUCKET"), objectName, expiry, reqParams)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, serializer.Response{
// 			Code: 50002,
// 			Msg: "OSS 配置错误，下载 url",
// 			Error: err,
// 		})

// 		return 
// 	}

// 	c.JSON(200, serializer.Response{
// 		Data: map[string]interface{}{
// 			"key": objectName,
// 			"put": signedPutURL,  // 上传地址
// 			"get": signedGetURL,  // 下载地址
// 		},
// 	})
// }


func UploadFile() (string, bool) {

	objectName := uuid.Must(uuid.NewRandom()).String() + ".jpg"
	filePath := "/upload/avatar/" + objectName

	_, err := MinioClient.FPutObject(os.Getenv("MINIO_BUCKET"), objectName, filePath, minio.PutObjectOptions{ContentType: "image/jpeg"})
	
	if err != nil {
		util.Log().Println(err.Error())
		return objectName, false
	}
	return objectName, true
}


func GetFileURL(filename string) string {
	expiry := time.Second * 24 * 60 * 60 // 1 day
	reqParams := make(url.Values)

	signedGetURL, err := MinioClient.PresignedGetObject(os.Getenv("MINIO_BUCKET"), filename, expiry, reqParams)
	if err != nil {
		util.Log().Println(err.Error())
		return ""
	}

	return fmt.Sprintf("%s", signedGetURL)

}