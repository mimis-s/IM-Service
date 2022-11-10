package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/account"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/chat"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/IM-Service/src/services/gateway"
	"github.com/mimis-s/IM-Service/src/services/home"
	"github.com/mimis-s/IM-Service/src/services/home/api_home"
	"github.com/mimis-s/IM-Service/src/services/relay"
	"github.com/mimis-s/IM-Service/web_client"
)

func initRpcxClient() {
	etcdAddrs := []string{}
	var timeout time.Duration
	etcdBasePath := ""
	isLocal := true
	api_home.SingleNewHomeClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_chat.SingleNewChatClient(etcdAddrs, timeout, etcdBasePath, isLocal)
	api_account.SingleNewAccountClient(etcdAddrs, timeout, etcdBasePath, isLocal)
}

func main() {

	initRpcxClient()

	ctx, cancel := context.WithCancel(context.Background())

	// 初始化消息队列生产者
	url := "amqp://dev:dev123@localhost:5672/"
	durable := false
	err := event.InitProducers(url, durable)
	if err != nil {
		panic(err)
	}

	// 启动每个服务
	go gateway.Boot(ctx)
	go home.Boot(ctx)
	go chat.Boot(ctx)
	go account.Boot(ctx)
	go relay.Boot(ctx)

	// 运行网页客户端
	go web_client.Boot(ctx)

	go GracefulStop(cancel)
	select {
	case <-ctx.Done():
		time.Sleep(time.Second * 5)
		fmt.Printf("stop service:%v", "im-service")
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
