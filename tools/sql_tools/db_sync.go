package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"xorm.io/core"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

/*
	将golang结构转为数据库结构
*/

func main() {
	user := flag.String("u", "root", "mysql user")
	pwd := flag.String("p", "dev123", "mysql password")
	addr := flag.String("a", "localhost", "mysql address")
	database := flag.String("d", "im_zhangbin", "mysql database")
	cmd := flag.String("cmd", "sync", "sync|clean|drop_db|drop_shard|shard")

	switch *cmd {
	case "sync":
		orm := getEngine(*user, *pwd, *addr, *database, "")

		err := dbmodel.InitSync(orm)
		if err != nil {
			fmt.Printf("init database[%v] error:%v\n", *database, err)
			os.Exit(1)
		}
		fmt.Printf("sync database[%v] ok\n", *database)
	case "drop_db": // 删除数据库
		orm := getEngine(*user, *pwd, *addr, "", "")

		// 删除数据库
		delDatabaseSql := fmt.Sprintf("drop database %v", *database)
		_, err := orm.Exec(delDatabaseSql)
		if err != nil {
			fmt.Printf("del database[%v] is error:%v", *database, err)
			os.Exit(1)
		}
		fmt.Printf("drop database[%v] ok\n", *database)
	case "clean": // 清理数据库
		// 删除所有分表数据
		dbTables := dbmodel.GetShardInfo()
		for _, table := range dbTables {
			for index := 0; index < table.SubTableNum(); index++ {
				if dbmodel.ShardDatabaseName()[string(table.BindSubTreasury())] {
					tableName := table.SubName()
					if index > 0 {
						tableName += "_" + strconv.Itoa(index)
					}
					ormTemp := getEngine(*user, *pwd, *addr, *database, string(table.BindSubTreasury()))
					truncateSql := fmt.Sprintf("truncate table %v", tableName)
					_, err := ormTemp.Exec(truncateSql)
					if err != nil {
						fmt.Printf("shard database[%v] table[%v] truncate is err:[%v]", dbmodel.ShardDatabaseName(),
							tableName, err)
						os.Exit(1)
					}
					fmt.Printf("truncate database[%v] table[%v] is ok\n", table.BindSubTreasury(), tableName)

				} else {
					fmt.Printf("shard database[%v] not find table[%v] bind name[%v]", dbmodel.ShardDatabaseName(),
						table.SubName(), table.BindSubTreasury())
					os.Exit(1)
				}
			}
		}
	case "drop_shard": // 删除所有数据库
		orm := getEngine(*user, *pwd, *addr, "", "")

		// 删除数据库
		for name, _ := range dbmodel.ShardDatabaseName() {
			delDatabaseSql := fmt.Sprintf("drop database %v", (*database)+"_"+name)
			_, err := orm.Exec(delDatabaseSql)
			if err != nil {
				fmt.Printf("del database[%v] is error:%v", (*database)+"_"+name, err)
			}
			fmt.Printf("drop database[%v] is ok\n", (*database)+"_"+name)
		}
	case "shard": // 创建数据库, 分库分表
		orm := getEngine(*user, *pwd, *addr, "", "")

		// 删除数据库
		for name, _ := range dbmodel.ShardDatabaseName() {
			createDatabaseSql := fmt.Sprintf("drop database %v", (*database)+"_"+name)
			_, err := orm.Exec(createDatabaseSql)
			if err != nil {
				fmt.Printf("del database[%v] is error:%v", (*database)+"_"+name, err)
			}
			fmt.Printf("drop database[%v] is ok\n", (*database)+"_"+name)
		}

		// 先创建分库
		for name, _ := range dbmodel.ShardDatabaseName() {
			createDatabaseSql := fmt.Sprintf("create database %v", (*database)+"_"+name)
			_, err := orm.Exec(createDatabaseSql)
			if err != nil {
				fmt.Printf("create database[%v] is error:%v", (*database)+"_"+name, err)
				os.Exit(1)
			}
			fmt.Printf("create database[%v] is ok\n", (*database)+"_"+name)
		}

		// 分表
		dbTables := dbmodel.GetShardInfo()
		for _, table := range dbTables {
			for index := 0; index < table.SubTableNum(); index++ {
				if dbmodel.ShardDatabaseName()[string(table.BindSubTreasury())] {
					suffixName := ""
					if index > 0 {
						suffixName = "_" + strconv.Itoa(index)
					}
					ormTemp := getEngine(*user, *pwd, *addr, *database, string(table.BindSubTreasury()))
					tbMapper := core.NewSuffixMapper(core.GonicMapper{}, suffixName)
					ormTemp.SetTableMapper(tbMapper)
					tableName := table.SubName() + suffixName

					err := ormTemp.Sync2(table)
					if err != nil {
						fmt.Printf("shard database[%v] table[%v] create is err:[%v]", dbmodel.ShardDatabaseName(),
							tableName, err)
						os.Exit(1)
					}
					fmt.Printf("create database[%v] table[%v] is ok\n", table.BindSubTreasury(), tableName)

				} else {
					fmt.Printf("shard database[%v] not find table[%v] bind name[%v]", dbmodel.ShardDatabaseName(),
						table.SubName(), table.BindSubTreasury())
					os.Exit(1)
				}
			}
		}
	}

}

var mapDatabaseEngine = make(map[string]*xorm.Engine)

func getEngine(user, pwd, addr, databasePrefix, dataBaseName string) *xorm.Engine {
	name := databasePrefix
	if dataBaseName != "" {
		name = databasePrefix + "_" + dataBaseName
	}

	if mapDatabaseEngine[name] == nil {

		dataSourceName := user + ":" + pwd + "@tcp(" + addr + ")/" + name + "?charset=utf8"

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

		mapDatabaseEngine[name] = orm
	}
	return mapDatabaseEngine[name]
}
