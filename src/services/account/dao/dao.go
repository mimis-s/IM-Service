package dao

import (
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mimis-s/IM-Service/src/common/dbmodel"
	"github.com/mimis-s/golang_tools/dfs"
	"xorm.io/xorm"
)

type TableSession struct {
	Account *xorm.Session
}

type Dao struct {
	Session    *TableSession
	Cache      *redis.Client
	DFSHandler dfs.DFSHandler
}

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
		Password:           "",
		DB:                 1,
		MaxRetries:         2,
		DialTimeout:        time.Second * 10,
		ReadTimeout:        time.Second * 5,
		WriteTimeout:       time.Second * 5,
		PoolTimeout:        time.Second * 10,
		IdleTimeout:        time.Minute * 10,
		IdleCheckFrequency: time.Second * 30,
	})
	return client
}

func New() (*Dao, error) {

	// 初始化数据库xorm
	engine, err := xorm.NewEngine("mysql", "root:dev123@tcp(localhost:3306)/im_zhangbin?charset=utf8")
	if err != nil {
		panic(err)
	}

	// 分布式文件存储
	dfsConfig := &dfs.Config{
		Type:       dfs.DFSType_Minio,
		Bucket:     "im_zhangbin",
		ExpireDays: 30,
		Url:        "localhost:9000",
		KeyID:      "admin",
		Key:        "admin123456",
	}
	dfsHandler, err := dfs.NewDFSHandler(dfsConfig)
	if err != nil {
		panic(err)
	}

	// 新建桶
	err = dfsHandler.TryMakeBucket()
	if err != nil {
		panic(err)
	}

	dao := &Dao{
		Session: &TableSession{
			Account: engine.Table((*dbmodel.AccountUser).SubName(nil)),
		},
		Cache:      newRedisClient(),
		DFSHandler: dfsHandler,
	}

	return dao, nil
}
