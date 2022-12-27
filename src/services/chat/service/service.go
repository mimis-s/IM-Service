package service

import (
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/IM-Service/src/services/chat/dao"
)

var S *Service

type Service struct {
	Dao *dao.Dao
}

func Init(configOptions *boot_config.ConfigOptions) *Service {
	d, err := dao.New(configOptions)
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
	api_chat.NewChatServiceAndRun(listenAddr, addr, etcdAddrs, S, etcdBasePath, isLocal)

	return S
}
