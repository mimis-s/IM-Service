package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/services/main/api_main"
)

// 处理分发客户端发送的消息
func (s *Service) ClientRequestHandleProto(ctx context.Context, req *api_main.ClientRequestHandleReq, res *api_main.ClientRequestHandleRes) error {
	return nil
}

func (s *Service) ClientRequestHandleJson(ctx context.Context, req *api_main.ClientRequestHandleReq, res *api_main.ClientRequestHandleRes) error {
	return nil
}
