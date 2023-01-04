package service

import (
	"context"

	"github.com/mimis-s/IM-Service/src/services/message/api_message"
)

// 存储离线消息
func (s *Service) SaveOffLineMessage(ctx context.Context, req *api_message.SaveOffLineMessageReq, res *api_message.SaveOffLineMessageRes) error {

	err = s.Dao.AddUserOneOfflineMessage(req.Data.SenderID, req.Data.ReceiverID, req.Data)
	if err != nil {

	}
	return nil
}
