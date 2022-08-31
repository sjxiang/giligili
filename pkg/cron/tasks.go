package cron

import "giligili/pkg/cache"

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.Redis.Client.Del(cache.Redis.Context, cache.DailyRankKey).Err()
}