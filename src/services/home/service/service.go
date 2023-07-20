package service

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/account"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/chat"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/friends"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/message"
	rpcxService "github.com/mimis-s/golang_tools/rpcx/service"
)

var ProviderSet = wire.NewSet(NewServiceHandler)

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
}

func NewServiceHandler(rpcSvc *rpcxService.ServerManage) (*Service, error) {

	S = &Service{}

	// 绑定rpcx服务
	err := api_home.RegisterHomeService(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	return S, nil
}
