package service

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/commonproto/im_error_proto"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/golang_tools/lib"
)

const (
	maxGetChatHistoryNum = 20 // 一次获取的最大历史记录条数
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

// 获取私聊历史记录
func (s *Service) GetSingleChatHistory(ctx context.Context, req *api_message.GetSingleChatHistoryReq, res *api_message.GetSingleChatHistoryRes) error {

	res.Data = &im_home_proto.GetSingleChatHistoryRes{
		FriendID: req.Data.FriendID,
		Data:     make([]*im_home_proto.ChatMessage, 0),
	}

	var maxMessageID int64
	var minMessageID int64
	var err error
	if req.Data.MaxNotGainMessageID == 0 {
		maxMessageID, err = s.Dao.GetHistoryMessageID(req.ClientInfo.UserID, req.Data.FriendID)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get friend[%v] history, but get history message id is err:%v",
				req.ClientInfo.UserID, req.Data.FriendID, err)
			im_log.Warn(errStr)
			res.ErrCode = im_error_proto.ErrCode_db_read_err
			return fmt.Errorf(errStr)
		}
		if maxMessageID == -1 {
			im_log.Info("user[%v] get friend[%v] history, but not find", req.ClientInfo.UserID, req.Data.FriendID)
			return nil
		}
		minMessageID = lib.MaxInt64(0, maxMessageID-maxGetChatHistoryNum+1)
	}

	historyMessages, err := s.Dao.GetHistoryMessage(req.ClientInfo.UserID, req.Data.FriendID, minMessageID, maxMessageID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friend[%v] history [%v]-[%v], but get history message db is err:%v",
			req.ClientInfo.UserID, req.Data.FriendID, minMessageID, maxMessageID, err)
		im_log.Warn(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		return fmt.Errorf(errStr)
	}
	for _, m := range historyMessages {
		res.Data.Data = append(res.Data.Data, m.MessageData.HistoryData)
	}

	return nil
}

// 读取离线消息
func (s *Service) ReadOfflineMessage(ctx context.Context, req *api_message.ReadOfflineMessageReq, res *api_message.ReadOfflineMessageRes) error {

	// 读取离线消息并且把离线消息转换成历史消息
	offlineMessage, err := s.Dao.GetUserOfflineMessage(req.ClientInfo.UserID, req.Data.FriendID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get friend[%v] off line message is err:%v",
			req.ClientInfo.UserID, req.Data.FriendID, err)
		im_log.Warn(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		return fmt.Errorf(errStr)
	}

	for _, offline := range offlineMessage {

		// 插入历史数据
		offline.MessageStatus = im_home_proto.MessageStatus_Enum_EnumRead
		err = s.Dao.AddHistoryMessage(offline.SenderID, offline.ReceiverID, offline)
		if err != nil {
			res.ErrCode = im_error_proto.ErrCode_db_write_err
			errStr := fmt.Sprintf("user[%v] add history message[%v] is err:%v",
				req.ClientInfo.UserID, req.Data, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}

		// 删除离线数据
		err = s.Dao.DelUserOneOfflineMessage(offline.SenderID, offline.ReceiverID)
		if err != nil {
			errStr := fmt.Sprintf("user[%v] del friend[%v] off line message is err:%v",
				req.ClientInfo.UserID, req.Data.FriendID, err)
			im_log.Warn(errStr)
			res.ErrCode = im_error_proto.ErrCode_db_write_err
			return fmt.Errorf(errStr)
		}

	}
	res.Data = &im_home_proto.ReadOfflineMessageRes{
		Data: offlineMessage,
	}

	return nil
}
