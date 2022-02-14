package main

import (
	"IM-Service/src/services/lobby"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 启动每个服务

	go lobby.Boot(ctx)

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
