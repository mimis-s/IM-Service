package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
)

func (s *Service) ChatSingle(ctx context.Context, req *api_chat.ChatSingleReq, res *api_chat.ChatSingleRes) error {
	res.Data = req.Data
	return nil
}
