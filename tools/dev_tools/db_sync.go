package main

import (
	"flag"
	"fmt"
	"os"

	"IM-Service/dbmodel"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func main() {
	user := flag.String("u", "zhangbin", "mysql user")
	pwd := flag.String("p", "zb1998810", "mysql password")
	addr := flag.String("a", "139.155.88.221", "mysql address")
	database := flag.String("d", "im_zhangbin", "mysql database")
	cmd := flag.String("cmd", "sync", "sync|clean")

	dataSourceName := *user + ":" + *pwd + "@tcp(" + *addr + ")/" + *database + "?charset=utf8"

	// dialects.RegisterDriver("mysql", &mysqlDriver{})
	// 修复了sync时，结构体包含json的字段时，sync到库里的字段是text类型
	// dialects.RegisterDialect("mysql", func() dialects.Dialect { return &dialect.Mysql{} })

	orm, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("new xorm engine[%v] error[%v]\n", dataSourceName, err)
		os.Exit(1)
	}

	orm.SetMaxOpenConns(64)
	orm.SetMaxIdleConns(64)

	err = orm.Ping()
	if err != nil {
		fmt.Printf("ping database[%v] error[%v]\n", dataSourceName, err)
		os.Exit(1)
	}

	orm.ShowSQL(true)
	orm.SetLogLevel(log.LOG_WARNING)

	switch *cmd {
	case "sync":
		err = dbmodel.InitDbSync(orm)
		if err != nil {
			fmt.Printf("init database[%v] error:%v\n", dataSourceName, err)
			os.Exit(1)
		}
		fmt.Printf("sync database[%v] ok\n", dataSourceName)
	case "clean": // 清理数据库
		// err = dbmodel.TruncateAll(orm)
		// if err != nil {
		// 	fmt.Printf("truncate database[%v] error:%v\n", dataSourceName, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("truncate database[%v] ok\n", dataSourceName)
	}

}
