package service

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/IM-Service/src/services/message/dao"
	rpcxService "github.com/mimis-s/golang_tools/rpcx/service"
)

var ProviderSet = wire.NewSet(dao.ProviderSet, NewServiceHandler)

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao *dao.Dao
}

func NewServiceHandler(rpcSvc *rpcxService.ServerManage, d *dao.Dao) (*Service, error) {

	S = &Service{
		Dao: d,
	}

	// 绑定rpcx服务
	err := api_message.RegisterMessageService(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	return S, nil
}
