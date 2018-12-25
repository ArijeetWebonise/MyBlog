package redis

import (
	"github.com/go-redis/redis"
)

type RedisService interface {
	Init() (string, error)
}

type RedisServiceImpl struct {
	Addr     string
	Password string
	DB       int

	client *redis.Client
}

func (rsi *RedisServiceImpl) Init() (string, error) {
	options := &redis.Options{
		Addr:     rsi.Addr,
		Password: rsi.Password,
		DB:       rsi.DB,
	}
	rsi.client = redis.NewClient(options)

	return rsi.client.Ping().Result()
}
