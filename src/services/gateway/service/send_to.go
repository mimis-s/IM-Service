package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/gateway/api_gateway"
	"github.com/mimis-s/golang_tools/net/clientConn"
)

/*
	主要用于给对应userID的客户端发送消息
*/

func (s *Service) SendToClient(ctx context.Context, req *api_gateway.SendToClientReq, res *api_gateway.SendToClientRes) error {
	c, ok := cacheClient.Load(req.ReceiverID)
	if !ok {
		// 接收方未登录, 消息无法到达
		errStr := fmt.Sprintf("user[%v] send msg[%v] to user[%v], but receiver not online",
			req.SenderID, req.MsgTag, req.ReceiverID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	session := c.(*Session)

	resClientMsg := &clientConn.ClientMsg{
		Tag: int(req.MsgTag),
		Msg: []byte(req.Payload),
	}

	err := session.GetClientConn().SendMsg(resClientMsg)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] send msg[%v] to user[%v], but is err:%v",
			req.SenderID, req.MsgTag, req.ReceiverID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	im_log.Info("user[%v] send to user[%v] msg_id[%v]", req.SenderID, req.ReceiverID, req.MsgTag)
	return nil
}

func (s *Service) NotifyClient(ctx context.Context, req *api_gateway.NotifyClientReq, res *api_gateway.NotifyClientRes) error {
	c, ok := cacheClient.Load(req.UserID)
	if !ok {
		// 接收方未登录, 消息无法到达
		errStr := fmt.Sprintf("notify send msg id[%v] to user[%v], but receiver not online",
			req.MsgTag, req.UserID)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	session := c.(*Session)

	resClientMsg := &clientConn.ClientMsg{
		Tag: int(req.MsgTag),
		Msg: []byte(req.Payload),
	}

	err := session.GetClientConn().SendMsg(resClientMsg)
	if err != nil {
		errStr := fmt.Sprintf("notify send msg id[%v] to user[%v], but send id err:%v",
			req.MsgTag, req.UserID, err)
		fmt.Println(errStr)
		return fmt.Errorf(errStr)
	}

	im_log.Info("notify send to user[%v] msg_id[%v]", req.UserID, req.MsgTag)
	return nil
}
