package account

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/service"
)

func Boot(ctx context.Context) {
	service.Init()

	select {
	case <-ctx.Done():
		im_log.Info("gateway service is stop\n")
	}
}
