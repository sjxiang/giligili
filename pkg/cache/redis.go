package cache

import (
	"context"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"

	"giligili/pkg/util"
)

type RedisClient struct {
	Client *redis.Client
	Context context.Context
}

// 确保全局 Redis 对象，只实例一次
var once sync.Once

// 全局实例
var Redis *RedisClient

// ConnectRedis 连接 redis，设置全局 Redis 对象
func ConnectRedis(address string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, password, db)
	})
}


// NewClient 创建 1 个新的 redis 连接
func NewClient(address string, password string, db int) *RedisClient {
	

	// 初始化自定义的 RedisClient 实例
	rds := &RedisClient{}

	// 使用默认的 context
	rds.Context = context.Background()
	
	// 使用 redis 库里的[redis/v8] NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr: 		address,
		Password: 	password,
		DB: 		db,
	})

	// 测试连接
	err := rds.Ping()
	if err != nil {
		util.Log().Panic("连接 Redis 不成功, ", err)
	}

	return rds
}


// Ping 测试 redis 连接是否正常
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}


// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		util.Log().Error("Redis", "Get", err.Error())
		return false
	}

	return true
}


// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			util.Log().Error("Redis", "Get", err.Error())
		}
		return ""
	} 

	return result
}


// Increment 
// 当参数数量为 1 时，默认其值 +1
// 当参数数量为 2 时，参数 1 为 key，参数 2 为要增加的值，int64 类型
func (rds RedisClient) Increment(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			util.Log().Error("Redis", "Incr", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			util.Log().Error("Redis", "IncrBy", err.Error())
			return false
		}
	default:
		util.Log().Error("Redis", "Incr", "参数过多")
		return false
	}

	return true
}



// ZIncrementBy
func (rds RedisClient) ZIncrementBy(key string, value float64, member string) bool {
	if err := rds.Client.ZIncrBy(rds.Context, key, value, member).Err(); err != nil {
		util.Log().Error("Redis", "ZIncrBy", err.Error())
		return false		
	}
	return true
}



