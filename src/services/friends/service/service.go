package service

import (
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/IM-Service/src/services/friends/dao"
)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func Init() *Service {
	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao: d,
	}

	listenAddr := ""
	addr := ""
	etcdAddrs := []string{}
	etcdBasePath := ""
	isLocal := true
	// 启动rpcx服务
	api_friends.NewFriendsServiceAndRun(listenAddr, addr, etcdAddrs, S, etcdBasePath, isLocal)
	return S
}
