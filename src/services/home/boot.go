package home

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/home/service"
)

func Boot(ctx context.Context) {
	service.Init()

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
