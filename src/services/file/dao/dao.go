package dao

import (
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"xorm.io/xorm"
)

type Dao struct {
	db *xorm.Engine
}

func New(configOptions *boot_config.ConfigOptions) (*Dao, error) {
	// engine, err := common_client.NewEngine(common_client.ENUM_MYSQL_DB_TAG_Chat)
	// if err != nil {
	// 	im_log.Warn("chat dao new engine is err:%v", err)
	// 	return nil, fmt.Errorf("chat dao new engine is err:%v", err)
	// }

	// dao := &Dao{
	// 	db: engine,
	// }

	return nil, nil
}
