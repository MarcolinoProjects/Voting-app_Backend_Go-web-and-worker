package config

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Client *redis.Client
}

func (c *RedisConfig) shutdown() {
	err := c.Client.Close()
	failOnError(err, "Failed to close connection with Redis")
}

func (c *RedisConfig) initialize() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_HOST", "localhost:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	c.Client = rdb
}
