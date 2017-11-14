package pub

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type RedisOpt struct {
	Address    string
	DB         string
	Password   string
	MaxRetries int
}

func (r *RedisOpt) Connect() *redis.Client {
	db, _ := strconv.Atoi(r.DB)
	return redis.NewClient(&redis.Options{
		Addr:         r.Address,
		Password:     r.Password,
		DB:           db,
		WriteTimeout: time.Second * 10,
		MaxRetries:   r.MaxRetries,
	})
}
