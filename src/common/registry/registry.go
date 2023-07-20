package registry

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/IM-Service/src/services/overrall/api_overrall"
	"github.com/mimis-s/IM-Service/web_client"
	"github.com/mimis-s/golang_tools/app"
	"github.com/mimis-s/golang_tools/zlog"
)

// var (
// 	// 只带rpc服务器的app
// 	DefaultSvcAppProviderSet = wire.NewSet(NewService)
// 	// DefaultAppProviderSet            = wire.NewSet()
// 	DefaultAppPeerServiceProviderSet = wire.NewSet(NewPeer2PeerService)
// )

func NewDefRegistry() *app.Registry {
	s := app.NewRegistry(
		app.AddRegistryBootConfigFile(boot_config.BootConfigData),
		app.AddRegistryExBootFlags(boot_config.CustomBootFlagsData),
		app.AddRegistryGlobalCmdFlag(boot_config.BootFlagsData),
	)

	s.AddInitTask("初始化rpc调用客户端", func() error {

		// 初始化rpc调用代码
		etcdAddrs := boot_config.BootConfigData.Etcd.Addrs
		timeout := time.Duration(boot_config.BootConfigData.Etcd.Timeout * int(time.Second))
		etcdBasePath := boot_config.BootConfigData.Etcd.EtcdBasePath
		isLocal := boot_config.BootConfigData.IsLocal
		api_home.SingleNewHomeClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_chat.SingleNewChatClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_account.SingleNewAccountClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_friends.SingleNewFriendsClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_overrall.SingleNewOverrallClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_message.SingleNewMessageClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		api_gateway.SingleNewGatewayClient(etcdAddrs, timeout, etcdBasePath, isLocal)

		// 日志
		zlog.NewLogger(boot_config.BootConfigData.Log.Path + "/" + "log.log")

		ctx, cancel := context.WithCancel(context.Background())

		// 初始化消息队列生产者
		// url := "amqp://dev:dev123@localhost:5672/"
		// durable := false
		err := event.InitProducers(boot_config.BootConfigData.MQ.Url, boot_config.BootConfigData.MQ.Durable)
		if err != nil {
			panic(err)
		}

		// 注册数据库
		common_client.RegisterParseMysql(boot_config.BootConfigData.DB)

		// 运行网页客户端
		go web_client.Boot(ctx)

		go GracefulStop(cancel)
		return nil
	})

	return s
}

func GracefulStop(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
