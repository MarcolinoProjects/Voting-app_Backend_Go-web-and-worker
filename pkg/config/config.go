package config

import (
	"log"
	"os"
)

var AppContext *AppConfig

type AppConfig struct {
	Production   bool
	RabbitConfig RabbitConfig
	RedisConfig  RedisConfig
}

func InitializeAppConfig(production bool) *AppConfig {
	rabbitmq := RabbitConfig{}
	rabbitmq.initialize()
	redisConfig := RedisConfig{}
	redisConfig.initialize()
	AppContext = &AppConfig{
		Production:   production,
		RabbitConfig: rabbitmq,
		RedisConfig:  redisConfig,
	}
	return AppContext
}

func ShutDown() {
	AppContext.RabbitConfig.shutdown()
	AppContext.RedisConfig.shutdown()
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
