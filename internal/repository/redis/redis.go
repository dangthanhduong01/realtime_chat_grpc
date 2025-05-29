package redis

import "github.com/redis/go-redis/v9"

type RedisClient struct {
	Client *redis.Client
}

func New(url string) *RedisClient {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
