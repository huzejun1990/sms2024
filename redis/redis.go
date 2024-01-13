// @Author huzejun 2024/1/13 15:58:00
package redis

import (
	"github.com/redis/go-redis/v9"
	"sms2024/config"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Secret.GetString("Redis.Addr"),
		Password: config.Secret.GetString("Redis.Password"),
		DB:       0,
	})
}
