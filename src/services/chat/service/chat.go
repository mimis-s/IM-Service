package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
)

// 服务器要做的是转发和回发
func (s *Service) ChatSingle(ctx context.Context, req *api_chat.ChatSingleReq, res *api_chat.ChatSingleRes) error {

	// 通过消息队列转发给对方
	singleMessage := &event.SingleMessage{
		UserInfo: req.ClientInfo,
		Message:  req.Data,
	}
	event.Publish(event.Event_SingleMessage, singleMessage)

	res.Data = req.Data
	return nil
}
