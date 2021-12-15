package redis

import (
	"context"

	"github.com/go-redis/redis"
)

func NewRedisClient(ctx context.Context, url string) (*redis.Client, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	opt.PoolSize = 10
	opt.MaxRetries = 2

	client := redis.NewClient(opt)
	client = client.WithContext(ctx)

	return client, nil
}
