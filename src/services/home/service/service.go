package service

import (
	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/services/gateway/dao"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/account"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/chat"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/friends"
	_ "github.com/mimis-s/IM-Service/src/services/home/service/message"
)

var S *Service

type Service struct {
	// 主界面大厅服务,主要用于处理,分发客户端的消息
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
	api_home.NewHomeServiceAndRun(listenAddr, addr, etcdAddrs, S, etcdBasePath, isLocal)

	return S
}
