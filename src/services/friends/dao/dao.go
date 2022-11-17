package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"xorm.io/xorm"
)

type TableSession struct {
	Friends *xorm.Session
}

type Dao struct {
	Session *TableSession
}

func New() (*Dao, error) {

	// 初始化数据库xorm

	engine, err := xorm.NewEngine("mysql", "root:dev123@tcp(localhost:3306)/im_zhangbin?charset=utf8")
	if err != nil {
		panic(err)
	}

	dao := &Dao{
		Session: &TableSession{
			Friends: engine.Table((*dbmodel.Friends).SubName(nil)),
		},
	}

	return dao, nil
}