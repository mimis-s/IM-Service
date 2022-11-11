package relay

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/relay/job"
	"github.com/mimis-s/IM-Service/src/services/relay/service"
)

func Boot(ctx context.Context) {
	s := service.Init()
	job.InitMQ(s)

	select {
	case <-ctx.Done():
		im_log.Info("gateway service is stop\n")
	}
}
