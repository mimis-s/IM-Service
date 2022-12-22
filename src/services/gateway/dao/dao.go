package dao

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"xorm.io/xorm"
)

type TableSession struct {
	Gateway *xorm.Session
}

type Dao struct {
	Session *TableSession
}

func New() (*Dao, error) {

	_, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	if err != nil {
		im_log.Warn("gateway dao new engine is err:%v", err)
		return nil, fmt.Errorf("gateway dao new engine is err:%v", err)
	}
	dao := &Dao{}

	return dao, nil
}
