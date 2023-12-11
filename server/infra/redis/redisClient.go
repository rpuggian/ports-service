package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	Find(id string, body interface{}) error
	Store(id string, body interface{}, expiration int) error
}

type redisClient struct {
	client *redis.Client
}

func NewClient(host, port, password string) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &redisClient{
		client,
	}, nil
}

func (r *redisClient) Find(id string, body interface{}) error {
	dataBytes, err := r.client.Get(id).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, body)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisClient) Store(id string, body interface{}, expiration int) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = r.client.Set(id, data, time.Duration(expiration)).Result()
	if err != nil {
		return err
	}
	return nil
}
