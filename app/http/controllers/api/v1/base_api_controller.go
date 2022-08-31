// 处理 v1 业务逻辑

package v1

import (
	"fmt"
	"giligili/app/model/video"
	"giligili/pkg/cache"
	"giligili/pkg/database"
	"giligili/pkg/serializer"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


type BaseAPIController struct {
	
}


func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Msg: "Pong",
	})
}


// 获取排行
func DailyRank(c *gin.Context) {
	
	var videos []video.Video

	// 从 redis 读取点击前十的视频
	vids, _ := cache.Redis.Client.ZRevRange(cache.Redis.Context, cache.DailyRankKey, 0, 9).Result()
	
	if len(vids) > 1 {
		
		// 排序 ORDER BY FIELD(id, 12, 10)
		// 过滤软删除的记录 WHERE `vidoes`.`deleted_at` IS NULL
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))

		err := database.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, serializer.Response{
				Code: 50000,
				Msg: "数据库连接错误",
				Error: err.Error(),
			})

			return
		}
	}

	
	c.JSON(http.StatusOK, serializer.Response{
		Data: serializer.BuildVideos(videos),
	})
	
}	