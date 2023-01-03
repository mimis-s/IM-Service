package common_client

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
)

func NewRedisClient(configOptions *boot_config.ConfigOptions) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:               configOptions.BootConfigFile.Redis.Addr,
		Password:           configOptions.BootConfigFile.Redis.Password,
		DB:                 configOptions.BootConfigFile.Redis.DB,
		MaxRetries:         2,
		DialTimeout:        time.Second * 10,
		ReadTimeout:        time.Second * 5,
		WriteTimeout:       time.Second * 5,
		PoolTimeout:        time.Second * 10,
		IdleTimeout:        time.Minute * 10,
		IdleCheckFrequency: time.Second * 30,
	})
	return client
}
