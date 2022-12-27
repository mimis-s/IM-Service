package service

import (
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/overrall/api_overrall"
	"github.com/mimis-s/IM-Service/src/services/overrall/view"
)

var S *Service

type Service struct {
	// Dao *dao.Dao
}

func Init(configOptions *boot_config.ConfigOptions) *Service {
	S = &Service{}

	// 生成雪花算法对象
	GenerateUserIDIdgenObject = view.NewIdgenObject()

	listenAddr := configOptions.CommandFlags.RpcListenPort
	addr := configOptions.CommandFlags.RpcExposePort
	etcdAddrs := configOptions.BootConfigFile.Etcd.Addrs
	etcdBasePath := configOptions.BootConfigFile.Etcd.EtcdBasePath
	isLocal := configOptions.BootConfigFile.IsLocal
	// 启动rpcx服务
	api_overrall.NewOverrallServiceAndRun(listenAddr, addr, etcdAddrs, S, etcdBasePath, isLocal)
	return S
}
