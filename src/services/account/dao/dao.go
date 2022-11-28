package dao

import (
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type TableSession struct {
	Account *xorm.Session
}

type Dao struct {
	Session *TableSession
	Cache   *redis.Client
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
	dao := &Dao{
		Session: &TableSession{
			// Account: engine.Table((*dbmodel.AccountUser).SubName(nil)),
			Account: engine.Table("account_user"),
		},
		Cache: newRedisClient(),
	}

	return dao, nil
}
