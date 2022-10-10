package gateway

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/gateway/service"
)

func Boot(ctx context.Context) {
	// Http服务, web客户端连接
	service.Init("localhost:8888", "localhost:8998")

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
