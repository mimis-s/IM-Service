package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/common_client"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/chat"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/IM-Service/src/services/friends"
	"github.com/mimis-s/IM-Service/src/services/friends/api_friends"
	"github.com/mimis-s/IM-Service/src/services/gateway"
	"github.com/mimis-s/IM-Service/src/services/home"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/overrall"
	"github.com/mimis-s/IM-Service/src/services/overrall/api_overrall"
	"github.com/mimis-s/IM-Service/src/services/relay"
	"github.com/mimis-s/IM-Service/web_client"
)

func initRpcxClient(configOptions *boot_config.ConfigOptions) {
	etcdAddrs := configOptions.BootConfigFile.Etcd.Addrs
	timeout := time.Duration(configOptions.BootConfigFile.Etcd.Timeout * int(time.Second))
	etcdBasePath := configOptions.BootConfigFile.Etcd.EtcdBasePath
	isLocal := configOptions.BootConfigFile.IsLocal
	api_home.SingleNewHomeClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_chat.SingleNewChatClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_account.SingleNewAccountClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_friends.SingleNewFriendsClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_overrall.SingleNewOverrallClient(etcdAddrs, timeout, etcdBasePath, isLocal)
}

func main() {

	configOptions := boot_config.ParseBootConfigOptions()
	initRpcxClient(configOptions)

	ctx, cancel := context.WithCancel(context.Background())

	// 初始化消息队列生产者
	// url := "amqp://dev:dev123@localhost:5672/"
	// durable := false
	err := event.InitProducers(configOptions.BootConfigFile.MQ.Url, configOptions.BootConfigFile.MQ.Durable)
	if err != nil {
		panic(err)
	}

	// 注册数据库
	common_client.RegisterParseMysql(configOptions.BootConfigFile.DB)

	// 启动每个服务
	go gateway.Boot(ctx, configOptions)
	go home.Boot(ctx, configOptions)
	go chat.Boot(ctx, configOptions)
	go account.Boot(ctx, configOptions)
	go relay.Boot(ctx, configOptions)
	go friends.Boot(ctx, configOptions)
	go overrall.Boot(ctx, configOptions)

	// 运行网页客户端
	go web_client.Boot(ctx, configOptions)

	go GracefulStop(cancel)
	select {
	case <-ctx.Done():
		time.Sleep(time.Second * 5)
		im_log.Info("stop service:%v", "im-service")
	}
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
