package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/chat/api_chat"
	"github.com/mimis-s/golang_tools/zlog"
)

// 服务器要做的是转发和回发
func (s *Service) ChatSingle(ctx context.Context, req *api_chat.ChatSingleReq, res *api_chat.ChatSingleRes) error {

	// 通过消息队列转发给对方
	singleMessage := &event.SingleMessage{
		UserInfo: req.ClientInfo,
		Message:  req.Data,
	}

	// 对于自己的id不能信任客户端
	singleMessage.Message.SenderID = req.ClientInfo.UserID

	err := event.Publish(event.Event_SingleMessage, singleMessage)
	if err != nil {
		zlog.Error("err:%v", err)
		res.ErrCode = im_error_proto.ErrCode_common_unexpected_err
		return err
	}

	res.Data = req.Data
	return nil
}
