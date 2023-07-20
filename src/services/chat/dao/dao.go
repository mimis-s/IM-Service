package dao

import (
	"fmt"

	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/golang_tools/zlog"
	"xorm.io/xorm"
)

var ProviderSet = wire.NewSet(New)

type Dao struct {
	db *xorm.Engine
}

func New() (*Dao, error) {
	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	if err != nil {
		zlog.Warn("chat dao new engine is err:%v", err)
		return nil, fmt.Errorf("chat dao new engine is err:%v", err)
	}

	dao := &Dao{
		db: engine,
	}

	return dao, nil
}
