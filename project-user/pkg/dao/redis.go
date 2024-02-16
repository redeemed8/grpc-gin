package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
	log.Println("redis init...")
}

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func (rc *RedisCache) Put(key, value string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}
