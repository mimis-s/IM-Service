package service

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
	"github.com/mimis-s/IM-Service/src/services/overrall/api_overrall"
	"github.com/mimis-s/IM-Service/src/services/overrall/view"
	rpcxService "github.com/mimis-s/golang_tools/rpcx/service"
)

var ProviderSet = wire.NewSet(NewServiceHandler)

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao *dao.Dao
}

func NewServiceHandler(rpcSvc *rpcxService.ServerManage) (*Service, error) {

	S = &Service{}

	// 生成雪花算法对象
	GenerateUserIDIdgenObject = view.NewIdgenObject()

	// 绑定rpcx服务
	err := api_overrall.RegisterOverrallService(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	return S, nil
}
