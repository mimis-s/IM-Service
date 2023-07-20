package dao

import (
	"fmt"

	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/golang_tools/dfs"
	"github.com/mimis-s/golang_tools/zlog"
	"xorm.io/xorm"
)

var ProviderSet = wire.NewSet(New)

type Dao struct {
	db         *xorm.Engine
	cache      *common_client.RedisClient
	dfsHandler dfs.DFSHandler
}

func New() (*Dao, error) {
	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Message)
	if err != nil {
		zlog.Warn("message dao new engine is err:%v", err)
		return nil, fmt.Errorf("message dao new engine is err:%v", err)
	}

	dfsHandler, err := dfs.NewDFSHandler(&boot_config.BootConfigData.DFS)
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
		cache:      common_client.NewRedisClient(),
		dfsHandler: dfsHandler,
	}

	return dao, nil
}
