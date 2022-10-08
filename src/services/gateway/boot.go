package gateway

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/gateway/service"
)

func Boot(ctx context.Context) {
	service.Init("localhost:8888")

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
