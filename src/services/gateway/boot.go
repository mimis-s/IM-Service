package gateway

import (
	"IM-Service/src/services/gateway/service"
	"context"
	"fmt"
)

func Boot(ctx context.Context) {
	service.Init("localhost:8888")

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
