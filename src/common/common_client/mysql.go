package common_client

import (
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/golang_tools/zlog"
	"xorm.io/xorm"
)

const (
	ENUM_MYSQL_DB_TAG_Account = "account"
	ENUM_MYSQL_DB_TAG_Chat    = "chat"
	ENUM_MYSQL_DB_TAG_Friends = "friends"
	ENUM_MYSQL_DB_TAG_Gateway = "gateway"
	ENUM_MYSQL_DB_TAG_Group   = "group"
	ENUM_MYSQL_DB_TAG_Home    = "home"
	ENUM_MYSQL_DB_TAG_Message = "message"
	ENUM_MYSQL_DB_TAG_Relay   = "relay"
)

var mapMysqlTag = map[string]bool{
	ENUM_MYSQL_DB_TAG_Account: true,
	ENUM_MYSQL_DB_TAG_Chat:    true,
	ENUM_MYSQL_DB_TAG_Friends: true,
	ENUM_MYSQL_DB_TAG_Gateway: true,
	ENUM_MYSQL_DB_TAG_Group:   true,
	ENUM_MYSQL_DB_TAG_Home:    true,
	ENUM_MYSQL_DB_TAG_Message: true,
	ENUM_MYSQL_DB_TAG_Relay:   true,
}

var mapMysql = make(map[string]boot_config.DBManageConfig)

func RegisterParseMysql(dbManageConfigs []boot_config.DBManageConfig) {
	for _, db := range dbManageConfigs {
		if mapMysqlTag[db.Tag] {
			mapMysql[db.Tag] = db
		}
	}
}

func NewEngine(tag string) (*xorm.Engine, error) {

	if !mapMysqlTag[tag] {
		zlog.Warn("map tag[%v] service tag[%v] db not define", mapMysqlTag, tag)
		return nil, fmt.Errorf("map tag[%v] service tag[%v] db not define", mapMysqlTag, tag)
	}

	// 初始化数据库xorm
	url := ""
	if dbConfig, ok := mapMysql[tag]; ok {
		url = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8",
			dbConfig.Master.User, dbConfig.Master.Password, dbConfig.Master.Addr, dbConfig.Master.DBName)
	}

	if url == "" {
		zlog.Warn("service tag[%v] db url is nil", tag)
		return nil, fmt.Errorf("service tag[%v] db url is nil", tag)
	}

	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		zlog.Warn("service tag[%v] url[%v] db New Engine is err:%v", tag, url, err)
		return nil, fmt.Errorf("service tag[%v] url[%v] db New Engine is err:%v", tag, url, err)
	}
	return engine, nil
}
