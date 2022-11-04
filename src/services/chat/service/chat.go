package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
)

// 服务器要做的是转发和回发
func (s *Service) ChatSingle(ctx context.Context, req *api_chat.ChatSingleReq, res *api_chat.ChatSingleRes) error {

	// 转发

	res.Data = req.Data
	return nil
}
