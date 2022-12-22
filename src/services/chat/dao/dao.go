package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"xorm.io/xorm"
)

type TableSession struct {
	Chat *xorm.Session
}

type Dao struct {
	Session *TableSession
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {
	_, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	if err != nil {
		im_log.Warn("friends dao new engine is err:%v", err)
		return nil, fmt.Errorf("friends dao new engine is err:%v", err)
	}

	dao := &Dao{
		Session: &TableSession{
			// Chat: engine.Table((*dbmodel).SubName(nil)),
		},
	}

	return dao, nil
}
