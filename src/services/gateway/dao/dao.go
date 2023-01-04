package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"xorm.io/xorm"
)

type Dao struct {
	db *xorm.Engine
}

func New() (*Dao, error) {

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Gateway)
	if err != nil {
		im_log.Warn("gateway dao new engine is err:%v", err)
		return nil, fmt.Errorf("gateway dao new engine is err:%v", err)
	}
	dao := &Dao{db: engine}

	return dao, nil
}
