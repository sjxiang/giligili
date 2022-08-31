package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)


// VideoViewKey 视频点击数数的 key
// 拼接出 view:video:10
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%s", strconv.Itoa(int(id)))
}
