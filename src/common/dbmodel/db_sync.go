package dbmodel

import (
	"fmt"

	"xorm.io/xorm"
)

func InitDbSync(orm *xorm.Engine) error {
	var err error
	err = orm.Sync2(&AccountAccount{})
	if err != nil {
		return fmt.Errorf("sync &AccountAccount error:%v", err)
	}
	return nil
}
