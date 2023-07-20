package dao

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Friends)
	if err != nil {
		zlog.Warn("friends dao new engine is err:%v", err)
		return nil, fmt.Errorf("friends dao new engine is err:%v", err)
	}
	dao := &Dao{
		db: engine,
	}

	return dao, nil
}
