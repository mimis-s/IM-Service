package service

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/golang_tools/net"
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
	err := api_gateway.RegisterGatewayService(rpcSvc, S)
	if err != nil {
		return nil, err
	}

	webAddr := boot_config.CustomBootFlagsData.IP + ":" + boot_config.CustomBootFlagsData.WebPort

	httpServer := net.InitServer(webAddr, "http", NewSession)

	go func() {
		err := httpServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	tcpAddr := boot_config.CustomBootFlagsData.IP + ":" + boot_config.CustomBootFlagsData.Port

	// 客户端连接的TCP连接
	tcpServer := net.InitServer(tcpAddr, "tcp", NewSession)

	go func() {
		err := tcpServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	return S, nil
}
