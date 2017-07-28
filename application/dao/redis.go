package Dao

import (
	"github.com/go-redis/redis"
	"github.com/NiciiA/AuthRest/config"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	return client
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     Config.RedisConnection,
		Password: Config.RedisPassword, // no password set
		DB:       Config.RedisDB,  // use default DB
	})
}