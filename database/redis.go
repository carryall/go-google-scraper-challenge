package database

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func SetupRedisPool() {
	RedisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial:      connectRedis,
	}
}

func GetRedisPool() *redis.Pool {
	return RedisPool
}

func connectRedis() (redis.Conn, error) {
	redisUrl, err := web.AppConfig.String("redisUrl")
	if err != nil {
		logs.Error("Redis URL not found: ", err)
	}

	return redis.DialURL(redisUrl)
}
