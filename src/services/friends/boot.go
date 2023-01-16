package friends

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/friends/job"
	"github.com/mimis-s/IM-Service/src/services/friends/service"
)

func Boot(ctx context.Context, configOptions *boot_config.ConfigOptions) {
	s := service.Init(configOptions)
	job.InitMQ(s, configOptions)

	select {
	case <-ctx.Done():
		im_log.Info("gateway service is stop\n")
	}
}
