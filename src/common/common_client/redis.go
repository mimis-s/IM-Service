package common_client

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/im_log"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(configOptions *boot_config.ConfigOptions) *RedisClient {
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
	return &RedisClient{
		Client: client,
	}
}

// redis事务
func (r *RedisClient) Transaction(key string, transactionFun func(pipe redis.Pipeliner) (interface{}, error)) (interface{}, error) {
	pipe := r.Client.TxPipeline()

	// // 设置过期时间
	// ok, err := pipe.Expire(context.Background(), key, expireTime).Result()
	// if err != nil {
	// 	errStr := fmt.Sprintf("redis set key[%v] expire time[%v] is err:%v", key, expireTime, err)
	// 	im_log.Warn(errStr)
	// 	return nil, fmt.Errorf(errStr)
	// }
	// if !ok {
	// 	errStr := fmt.Sprintf("redis set key[%v] expire time[%v] is not ok", key, expireTime)
	// 	im_log.Warn(errStr)
	// 	return nil, fmt.Errorf(errStr)
	// }

	result, err := transactionFun(pipe)
	if err != nil {
		errStr := fmt.Sprintf("redis set key[%v] is err:%v", key, err)
		im_log.Warn(errStr)
		return nil, fmt.Errorf(errStr)
	}

	_, err = pipe.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}
