package gateway

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/gateway/service"
)

func Boot(ctx context.Context) {
	// TCP服务, web客户端连接
	service.Init("localhost:8888", "localhost:8998")

	select {
	case <-ctx.Done():
		im_log.Info("gateway service is stop\n")
	}
}
