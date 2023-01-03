package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"xorm.io/xorm"
)

type TableSession struct {
	HistoryMessage *xorm.Session
}

type Dao struct {
	Session *TableSession
	Cache   *common_client.RedisClient
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {
	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	if err != nil {
		im_log.Warn("message dao new engine is err:%v", err)
		return nil, fmt.Errorf("message dao new engine is err:%v", err)
	}

	dao := &Dao{
		Session: &TableSession{
			HistoryMessage: engine.Table((*dbmodel.HistoryMessage).SubName(nil)),
		},
		Cache: common_client.NewRedisClient(configOptions),
	}

	return dao, nil
}
