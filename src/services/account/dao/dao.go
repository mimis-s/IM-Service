package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Dao struct {
	Engine *xorm.Engine
}

func New() (*Dao, error) {

	// 初始化数据库xorm

	engine, err := xorm.NewEngine("mysql", "root:dev123@tcp(localhost:3306)/im_zhangbin?charset=utf8")
	if err != nil {
		panic(err)
	}
	dao := &Dao{
		Engine: engine,
	}

	return dao, nil
}
