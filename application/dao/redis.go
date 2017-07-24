package Dao

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	return client
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}