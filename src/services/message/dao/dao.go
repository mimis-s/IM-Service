package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/golang_tools/dfs"
	"xorm.io/xorm"
)

type Dao struct {
	db         *xorm.Engine
	cache      *common_client.RedisClient
	dfsHandler dfs.DFSHandler
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {
	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Message)
	if err != nil {
		im_log.Warn("message dao new engine is err:%v", err)
		return nil, fmt.Errorf("message dao new engine is err:%v", err)
	}

	dfsHandler, err := dfs.NewDFSHandler(&configOptions.BootConfigFile.DFS)
	if err != nil {
		panic(err)
	}

	// 新建桶
	err = dfsHandler.TryMakeBucket()
	if err != nil {
		panic(err)
	}

	dao := &Dao{
		db:         engine,
		cache:      common_client.NewRedisClient(configOptions),
		dfsHandler: dfsHandler,
	}

	return dao, nil
}
