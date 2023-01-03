package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/golang_tools/dfs"
	"xorm.io/xorm"
)

type TableSession struct {
	Account *xorm.Session
}

type Dao struct {
	Session    *TableSession
	Cache      *common_client.RedisClient
	DFSHandler dfs.DFSHandler
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {

	// 初始化数据库xorm

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Account)
	if err != nil {
		im_log.Warn("account dao new engine is err:%v", err)
		return nil, fmt.Errorf("account dao new engine is err:%v", err)
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
		Session: &TableSession{
			Account: engine.Table((*dbmodel.AccountUser).SubName(nil)),
		},
		Cache:      common_client.NewRedisClient(configOptions),
		DFSHandler: dfsHandler,
	}

	return dao, nil
}
