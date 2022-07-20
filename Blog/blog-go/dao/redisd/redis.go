package redisd

import (
	"github.com/go-redis/redis"
)

var RDB *redis.Client

//创建到redis的连接
func SetupRedisDb() error {

	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // use default DB
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
