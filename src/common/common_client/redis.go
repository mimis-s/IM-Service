package common_client

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:               boot_config.BootConfigData.Redis.Addr,
		Password:           boot_config.BootConfigData.Redis.Password,
		DB:                 boot_config.BootConfigData.Redis.DB,
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

// func (r *RedisClient) Transaction(key string, transactionFun func(pipe redis.Pipeliner) (interface{}, error)) (interface{}, error) {
// 	pipe := r.Client.TxPipeline()
// 	defer pipe.Close()
// 	// // 设置过期时间
// 	// ok, err := pipe.Expire(context.Background(), key, expireTime).Result()
// 	// if err != nil {
// 	// 	errStr := fmt.Sprintf("redis set key[%v] expire time[%v] is err:%v", key, expireTime, err)
// 	// 	zlog.Warn(errStr)
// 	// 	return nil, fmt.Errorf(errStr)
// 	// }
// 	// if !ok {
// 	// 	errStr := fmt.Sprintf("redis set key[%v] expire time[%v] is not ok", key, expireTime)
// 	// 	zlog.Warn(errStr)
// 	// 	return nil, fmt.Errorf(errStr)
// 	// }

// 	result, err := transactionFun(pipe)
// 	if err != nil {
// 		errStr := fmt.Sprintf("redis set key[%v] is err:%v", key, err)
// 		zlog.Warn(errStr)
// 		return result, fmt.Errorf(errStr)
// 	}

// 	_, err = pipe.Exec(context.Background())
// 	if err != nil {
// 		pipe.Discard()
// 		return result, err
// 	}
// 	return result, nil
// }
