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

	res.Data = req.Data

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

	dbChatData := &im_home_proto.ChatMessage{
		SenderID:         req.Data.SenderID,
		ReceiverID:       req.Data.ReceiverID,
		MessageID:        messageID,
		MessageFileInfos: make([]*im_home_proto.MessageFileRecap, 0, len(req.Data.MessageFileInfos)),
		SendTimeStamp:    req.Data.SendTimeStamp,
		MessageStatus:    req.Data.MessageStatus,
		Data:             req.Data.Data,
	}

	// 判断消息类型
	for _, fileMessage := range req.Data.MessageFileInfos {
		dbChatData.MessageFileInfos = append(dbChatData.MessageFileInfos, &im_home_proto.MessageFileRecap{
			FileName:        fileMessage.FileName,
			FileExtension:   fileMessage.FileExtension,
			FileSize:        fileMessage.FileSize,
			FileIndex:       fileMessage.FileIndex,
			MessageFileType: fileMessage.MessageFileType,
		})

		switch fileMessage.MessageFileType {
		case im_home_proto.MessageFileType_Enum_EnumImgType:
			err = s.Dao.UpLoadChatImage(req.Data.SenderID, req.Data.ReceiverID, messageID, int(fileMessage.FileIndex), []byte(fileMessage.FileData))
			if err != nil {
				errStr := fmt.Sprintf("user[%v] get user[%v][%v] upload message[%v] is err:%v", req.ClientInfo.UserID, req.Data.SenderID,
					req.Data.ReceiverID, messageID, err)
				im_log.Error(errStr)
				return fmt.Errorf(errStr)
			}
		case im_home_proto.MessageFileType_Enum_EnumFileType:
			err = s.Dao.UpLoadChatFile(req.Data.SenderID, req.Data.ReceiverID, messageID, int(fileMessage.FileIndex), []byte(fileMessage.FileData))
			if err != nil {
				errStr := fmt.Sprintf("user[%v] get user[%v][%v] upload message[%v] is err:%v", req.ClientInfo.UserID, req.Data.SenderID,
					req.Data.ReceiverID, messageID, err)
				im_log.Error(errStr)
				return fmt.Errorf(errStr)
			}
			fileMessage.FileData = ""
		default:
		}

	}

	if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Online {
		// 在线

		res.IsOnline = true
	} else if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Outline {
		// 离线
		err = s.Dao.AddUserOneUnReadMessage(req.Data.SenderID, req.Data.ReceiverID)
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

	// 写入数据库
	err = s.Dao.AddHistoryMessage(req.Data.SenderID, req.Data.ReceiverID, dbChatData)
	if err != nil {
		res.ErrCode = im_error_proto.ErrCode_db_write_err
		errStr := fmt.Sprintf("user[%v] add history message[%v] is err:%v",
			req.ClientInfo.UserID, req.Data, err)
		im_log.Error(errStr)
		return fmt.Errorf(errStr)
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
		for index, fileMessage := range m.MessageData.HistoryData.MessageFileInfos {

			// 判断消息类型
			switch fileMessage.MessageFileType {
			case im_home_proto.MessageFileType_Enum_EnumImgType:
				data, err := s.Dao.DownLoadChatImage(m.MessageData.HistoryData.SenderID, m.MessageData.HistoryData.ReceiverID, m.MessageData.HistoryData.MessageID, int(fileMessage.FileIndex))
				if err != nil {
					errStr := fmt.Sprintf("user[%v] get user[%v][%v] down load message[%v] is err:%v", req.ClientInfo.UserID, m.MessageData.HistoryData.SenderID,
						m.MessageData.HistoryData.ReceiverID, m.MessageData.HistoryData.MessageID, err)
					im_log.Error(errStr)
					return fmt.Errorf(errStr)
				}
				m.MessageData.HistoryData.MessageFileInfos[index].FileData = string(data)
			default:
			}
		}

		res.Data.Data = append(res.Data.Data, m.MessageData.HistoryData)
	}

	return nil
}

// 读取未读消息
func (s *Service) UnReadMessage(ctx context.Context, req *api_message.UnReadMessageReq, res *api_message.UnReadMessageRes) error {

	// 读取离线消息并且把离线消息转换成历史消息
	err := s.Dao.DelUserOneUnReadMessage(req.ClientInfo.UserID, req.Data.FriendID)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] del friend[%v] un read message is err:%v",
			req.ClientInfo.UserID, req.Data.FriendID, err)
		im_log.Warn(errStr)
		res.ErrCode = im_error_proto.ErrCode_db_read_err
		return fmt.Errorf(errStr)
	}

	res.Data = &im_home_proto.UnReadMessageRes{
		FriendID: req.Data.FriendID,
	}

	return nil
}

// 下载文件消息(文件或者图片)
func (s *Service) DownLoadFileMessage(ctx context.Context, req *api_message.DownLoadFileMessageReq, res *api_message.DownLoadFileMessageRes) error {

	fileData := ""
	switch req.Data.MessageFileType {
	case im_home_proto.MessageFileType_Enum_EnumImgType:
		data, err := s.Dao.DownLoadChatImage(req.ClientInfo.UserID, req.Data.FriendID,
			req.Data.MessageID, int(req.Data.FileIndex))
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get user[%v] down load img message[%v] is err:%v", req.ClientInfo.UserID,
				req.Data.FriendID, req.Data.MessageID, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
		fileData = string(data)
	case im_home_proto.MessageFileType_Enum_EnumFileType:
		data, err := s.Dao.DownLoadChatFile(req.ClientInfo.UserID, req.Data.FriendID,
			req.Data.MessageID, int(req.Data.FileIndex))
		if err != nil {
			errStr := fmt.Sprintf("user[%v] get user[%v] down load file message[%v] is err:%v", req.ClientInfo.UserID,
				req.Data.FriendID, req.Data.MessageID, err)
			im_log.Error(errStr)
			return fmt.Errorf(errStr)
		}
		fileData = string(data)
	default:
	}

	res.Data = &im_home_proto.DownLoadFileMessageRes{
		FriendID:        req.Data.FriendID,
		MessageID:       req.Data.MessageID,
		FileIndex:       req.Data.FileIndex,
		MessageFileType: req.Data.MessageFileType,
		FileData:        fileData,
	}

	return nil
}
