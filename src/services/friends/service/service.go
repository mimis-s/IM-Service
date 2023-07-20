package service

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/IM-Service/src/services/friends/dao"
	rpcxService "github.com/mimis-s/golang_tools/rpcx/service"
)

var ProviderSet = wire.NewSet(dao.ProviderSet, NewServiceHandler)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func NewServiceHandler(rpcSvc *rpcxService.ServerManage, d *dao.Dao) (*Service, error) {

	S = &Service{
		Dao: d,
	}

	// 绑定rpcx服务
	err := api_friends.RegisterFriendsService(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	return S, nil
}
