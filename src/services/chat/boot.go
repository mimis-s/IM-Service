// 聊天服务,不管是群聊还是私聊

package chat

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/services/chat/service"
)

func Boot(ctx context.Context) {
	service.Init()

	select {
	case <-ctx.Done():
		fmt.Printf("gateway service is stop\n")
	}
}
