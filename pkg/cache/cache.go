package cache

import (
	"context"
	"encoding/json"
	"github.com/cenkkoroglu/oz-fiber/pkg/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var client *redis.Client

func Init() error {
	config := config.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr: config.Cache.Addr,
		DB:   config.Cache.Db,
	})

	if _, err := Ping(); err != nil {
		return err
	}

	return nil
}

func Ping() (string, error) {
	pong, err := client.Ping(context.Background()).Result()
	return pong, err
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return client.Set(context.Background(), key, value, expiration).Err()
}

func SetEntity(key string, entity interface{}, expiration time.Duration) error {
	value, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return client.Set(context.Background(), key, value, expiration).Err()
}

func GetEntity(key string, resp interface{}) (err error) {
	value, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), resp)
}

func Exists(keys ...string) (int64, error) {
	return client.Exists(context.Background(), keys...).Result()
}

func Get(key string) (value string, err error) {
	return client.Get(context.Background(), key).Result()
}

func Del(key string) error {
	return client.Del(context.Background(), key).Err()
}
