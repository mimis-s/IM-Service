package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"xorm.io/xorm"
)

type TableSession struct {
	Friends *xorm.Session
}

type Dao struct {
	Session *TableSession
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Friends)
	if err != nil {
		im_log.Warn("friends dao new engine is err:%v", err)
		return nil, fmt.Errorf("friends dao new engine is err:%v", err)
	}
	dao := &Dao{
		Session: &TableSession{
			Friends: engine.Table((*dbmodel.Friends).SubName(nil)),
		},
	}

	return dao, nil
}
