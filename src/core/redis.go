package core

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var RedisMutex sync.Mutex

func RedisConnect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func RedisClose() {
	err := RedisClient.Close()
	if err != nil {
		log.Printf(err.Error())
	}
}

func GetRedisData(key string) string {
	RedisMutex.Lock()
	defer RedisMutex.Unlock()
	rs, err := RedisClient.Get(key).Result()
	if err != nil {

		return ""
	}
	return rs
}

func SetRedisData(key, data string, expiration time.Duration) {
	RedisMutex.Lock()
	defer RedisMutex.Unlock()
	RedisClient.Set(key, data, expiration)
}

func DeleteRedisData(key string) {
	RedisMutex.Lock()
	defer RedisMutex.Unlock()
	RedisClient.Del(key)
}
