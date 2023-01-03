package dao

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/im_log"
)

type Dao struct {
	Cache *redis.Client
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {
	_, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	if err != nil {
		im_log.Warn("message dao new engine is err:%v", err)
		return nil, fmt.Errorf("message dao new engine is err:%v", err)
	}

	dao := &Dao{
		Cache: common_client.NewRedisClient(configOptions),
	}

	return dao, nil
}
