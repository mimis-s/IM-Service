package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/golang_tools/dfs"
	"github.com/mimis-s/golang_tools/zlog"
	"xorm.io/xorm"
)

var ProviderSet = wire.NewSet(New)

type Dao struct {
	Db         *xorm.Engine
	cache      *common_client.RedisClient
	dfsHandler dfs.DFSHandler
}

func New() (*Dao, error) {

	// 初始化数据库xorm

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Account)
	if err != nil {
		zlog.Warn("account dao new engine is err:%v", err)
		return nil, fmt.Errorf("account dao new engine is err:%v", err)
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
		Db:         engine,
		cache:      common_client.NewRedisClient(),
		dfsHandler: dfsHandler,
	}

	// 清楚用户缓存
	dao.ClearCacheUser()

	return dao, nil
}
