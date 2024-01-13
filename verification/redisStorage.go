// @Author huzejun 2024/1/13 19:47:00
package verification

import (
	"context"
	v9 "github.com/redis/go-redis/v9"
	"sms2024/redis"
	"time"
)

type RedisStorage struct {
	client *v9.Client
}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{
		client: redis.Client,
	}
}

func (r *RedisStorage) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}
func (r *RedisStorage) Set(key string, val string, duration time.Duration) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, val, duration).Err()
}
func (r *RedisStorage) Del(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}
