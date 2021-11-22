package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	_ "github.com/go-sql-driver/mysql"
)

type TBJsonField1_AccountRoleInfo struct {
	RoleID int64 `json:"role_id,omitempty"`
	Zone   int   `json:"zone,omitempty"`
}

type TBJsonField_AccountRolesInfo struct {
	Roles []*TBJsonField1_AccountRoleInfo `json:"roles"`
}

type AccountAccount struct {
	Id        int                          `xorm:"not null pk autoincr INT(10)"`
	Account   string                       `xorm:"not null unique VARCHAR(128)"`
	Channel   string                       `xorm:"comment('渠道') VARCHAR(128)"`
	Platform  string                       `xorm:"comment('平台') VARCHAR(128)"`
	SdkId     int64                        `xorm:"comment('sdk唯一id') BIGINT(20)"`
	Country   string                       `xorm:"comment('国家') VARCHAR(128)"`
	Roles     TBJsonField_AccountRolesInfo `xorm:"comment('#TBJsonField_AccountRolesInfo#拥有角色') JSON"`
	BanReason int                          `xorm:"comment('封禁理由') INT(10)"`
	BanTime   int64                        `xorm:"comment('封禁到时') BIGINT(20)"`
	CreatedAt time.Time                    `xorm:"created"`
	UpdatedAt time.Time                    `xorm:"updated"`
	DeletedAt time.Time                    `xorm:"deleted"`
}

func main() {
	user := flag.String("u", "zhangbin", "mysql user")
	pwd := flag.String("p", "zb1998810", "mysql password")
	addr := flag.String("a", "139.155.88.221", "mysql address")
	database := flag.String("d", "IM-zhangbin", "mysql database")
	cmd := flag.String("cmd", "sync", "sync|clean")

	
	dataSourceName := *user + ":" + *pwd + "@tcp(" + *addr + ")/" + *database + "?charset=utf8"

	// dialects.RegisterDriver("mysql", &mysqlDriver{})
	// 修复了sync时，结构体包含json的字段时，sync到库里的字段是text类型
	// dialects.RegisterDialect("mysql", func() dialects.Dialect { return &dialect.Mysql{} })

	orm, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("new xorm engine[%v] error[%v]\n",dataSourceName, err)
		os.Exit(1)
	}

	orm.SetMaxOpenConns(64)
	orm.SetMaxIdleConns(64)

	err = orm.Ping()
	if err != nil {
		fmt.Printf("ping database[%v] error[%v]\n",dataSourceName, err)
		os.Exit(1)
	}

	orm.ShowSQL(true)
	orm.SetLogLevel(log.LOG_WARNING)

	switch *cmd {
	case "sync":
		err = orm.Sync2(&AccountAccount{})
		if err != nil {
			fmt.Errorf("sync &AccountAccount error:%v", err)
			os.Exit(1)
		}
		// err = dbmodel.InitSync(orm)
		// if err != nil {
		// 	fmt.Printf("init database[%v] error:%v\n", dataSourceName, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("sync database[%v] ok\n", dataSourceName)
	case "clean":
		// err = dbmodel.TruncateAll(orm)
		// if err != nil {
		// 	fmt.Printf("truncate database[%v] error:%v\n", dataSourceName, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("truncate database[%v] ok\n", dataSourceName)
	}

}