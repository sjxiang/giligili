package bootstrap

import (
	"os"
	"strconv"
	"giligili/pkg/cache"
)


// SetupRedis 初始化 redis 链接
func SetupRedis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)

	cache.ConnectRedis(
		os.Getenv("REDIS_ADDR"), 
		os.Getenv("REDIS_PW"), 
		int(db),
	)
}
