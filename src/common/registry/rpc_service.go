package registry

import (
	"github.com/google/wire"
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/golang_tools/rpcx/service"
)

var (
	DefaultAppRpcSet = wire.NewSet(NewRpcService)
)

func NewRpcService() (*service.ServerManage, error) {

	if boot_config.BootConfigData.IsLocal {
		return nil, nil
	}

	listenAddr := boot_config.CustomBootFlagsData.RpcListenPort
	exposeAddr := boot_config.CustomBootFlagsData.RpcExposePort
	etcdAddrs := boot_config.BootConfigData.Etcd.Addrs
	etcdBasePath := boot_config.BootConfigData.Etcd.EtcdBasePath

	s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath, listenAddr)
	if err != nil {
		return nil, err
	}
	return s, nil
}
