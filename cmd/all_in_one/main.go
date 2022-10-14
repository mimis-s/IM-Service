package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mimis-s/IM-Service/src/services/gateway"
	"github.com/mimis-s/IM-Service/src/services/home"
	"github.com/mimis-s/IM-Service/web_client"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	// 启动每个服务
	go gateway.Boot(ctx)
	go home.Boot(ctx)

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
