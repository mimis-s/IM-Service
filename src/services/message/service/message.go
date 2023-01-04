package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
)

// 存储私聊消息
func (s *Service) SaveSingleChatMessage(ctx context.Context, req *api_message.SaveSingleChatMessageReq, res *api_message.SaveSingleChatMessageRes) error {

	// 判断接收者是否在线
	getUserInfoReq := &api_account.GetUserInfoServiceReq{
		ClientInfo: req.ClientInfo,
		UserID:     req.Data.ReceiverID,
	}
	getUserInfoRes, err := api_account.GetUserInfoService(context.Background(), getUserInfoReq)
	if err != nil {
		res.ErrCode = getUserInfoRes.ErrCode
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but db is err:%v", req.ClientInfo.UserID, req.Data.ReceiverID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}

	messageID, err := s.Dao.GetMessageID(req.Data.SenderID, req.Data.ReceiverID)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		errStr := fmt.Sprintf("user[%v] get user[%v][%v] message id is err:%v", req.ClientInfo.UserID, req.Data.SenderID,
			req.Data.ReceiverID, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
	}
	req.Data.MessageID = messageID
	req.Data.MessageStatus = im_home_proto.MessageStatus_Enum_EnumArrive // 送达

	if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Online {
		// 在线
		err = s.Dao.AddHistoryMessage(req.Data.SenderID, req.Data.ReceiverID, req.Data)
		if err != nil {
			res.ErrCode = im_error_proto.ErrCode_db_write_err
			errStr := fmt.Sprintf("user[%v] add history message[%v] is err:%v",
				req.ClientInfo.UserID, req.Data, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
		res.IsOnline = true
	} else if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Outline {
		// 离线
		err = s.Dao.AddUserOneOfflineMessage(req.Data.SenderID, req.Data.ReceiverID, req.Data)
		if err != nil {
			res.ErrCode = im_error_proto.ErrCode_db_write_err
			errStr := fmt.Sprintf("user[%v] add off line message[%v] is err:%v",
				req.ClientInfo.UserID, req.Data, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
		res.IsOnline = false
	} else {
		im_log.Warn("user[%v] online status[%v] is not define", req.Data.ReceiverID, getUserInfoRes.Data.Status)
	}

	return nil
}
