package relay

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/relay/job"
	"github.com/mimis-s/IM-Service/src/services/relay/service"
)

func Boot(ctx context.Context) {
	s := service.Init()
	job.InitMQ(s)

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
