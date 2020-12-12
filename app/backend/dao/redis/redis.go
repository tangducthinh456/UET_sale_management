package redis

import (
	"SaleManagement/config"
	"github.com/go-redis/redis"
)

//RedisClient client to work with redis
var redisClient *redis.Client

//RedisClientInit connect to redis
func RedisClientInit() {
	rdConfig := config.GetRedisConfig()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     rdConfig.Host + ":6379",
		Password: rdConfig.Password,
		DB:       0,
	})
}

//GetRedisClient client for redis
func GetRedisClient() *redis.Client {
	return redisClient
}
