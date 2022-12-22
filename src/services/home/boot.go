package home

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/home/service"
)

func Boot(ctx context.Context, configOptions *boot_config.ConfigOptions) {
	service.Init()

	select {
	case <-ctx.Done():
		im_log.Info("gateway service is stop\n")
	}
}
