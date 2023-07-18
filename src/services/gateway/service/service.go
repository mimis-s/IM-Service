package service

import (
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/golang_tools/net"
)

var S *Service

// 现在的rpcx调用都不用,先使用本地调用
type Service struct {
	Dao *dao.Dao
}

func Init(configOptions *boot_config.ConfigOptions) *Service {

	d, err := dao.New()
	if err != nil {
		panic(err)
	}

	S = &Service{
		Dao: d,
	}

	listenAddr := configOptions.CommandFlags.RpcListenPort
	addr := configOptions.CommandFlags.RpcExposePort
	etcdAddrs := configOptions.BootConfigFile.Etcd.Addrs
	etcdBasePath := configOptions.BootConfigFile.Etcd.EtcdBasePath
	isLocal := configOptions.BootConfigFile.IsLocal
	// 启动rpcx服务
	api_gateway.NewGatewayServiceAndRun(listenAddr, addr, etcdAddrs, S, etcdBasePath, isLocal)

	webAddr := configOptions.CommandFlags.IP + ":" + configOptions.CommandFlags.WebPort

	httpServer := net.InitServer(webAddr, "http", NewSession)

	go func() {
		err := httpServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	tcpAddr := configOptions.CommandFlags.IP + ":" + configOptions.CommandFlags.Port

	// 客户端连接的TCP连接
	tcpServer := net.InitServer(tcpAddr, "tcp", NewSession)

	go func() {
		err := tcpServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	return S
}
