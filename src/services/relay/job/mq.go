package job

import (
	"context"
	"fmt"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/services/account/api_account"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/IM-Service/src/services/relay/relay_api"
	"github.com/mimis-s/IM-Service/src/services/relay/service"
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
	"github.com/mimis-s/golang_tools/zlog"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service) *Job {
	url := boot_config.BootConfigData.MQ.Url
	durable := boot_config.BootConfigData.MQ.Durable

	j := &Job{s}

	event.InitConsumers(url, durable,
		[]*rabbitmq.ConsumersQueue{
			{event.Event_SingleMessage, j.singleMessage},
			{event.Event_ApplyFriend, j.applyFriend},
			{event.Event_AgreeApplyFriend, j.agreeApplyFriend},
		},
	)

	return j
}

// 转发单人消息
func (j *Job) singleMessage(payload interface{}) error {
	singleMessage := payload.(*event.SingleMessage)

	// 存储消息
	saveSingleChatMessageReq := &api_message.SaveSingleChatMessageReq{
		ClientInfo: singleMessage.UserInfo,
		Data:       singleMessage.Message,
	}
	saveSingleChatMessageRes, err := api_message.SaveSingleChatMessage(context.Background(), saveSingleChatMessageReq)
	if err != nil {
		zlog.Warn("user[%v] to [%v] chat save is err:%v", singleMessage.Message.SenderID,
			singleMessage.Message.ReceiverID, err)
		return err
	}

	chatSingleToReceiver := &im_home_proto.ChatSingleToReceiver{
		Data: saveSingleChatMessageRes.Data,
	}
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.ChatSingleToReceiver{})

	getUserInfoServiceReq := &api_account.GetUserInfoServiceReq{
		ClientInfo: singleMessage.UserInfo,
		UserID:     singleMessage.UserInfo.UserID,
	}
	getUserInfoServiceRes, err := api_account.GetUserInfoService(context.Background(), getUserInfoServiceReq)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get self info is err:%v", singleMessage.UserInfo.UserID, err)
		zlog.Error(errStr)
		return fmt.Errorf(errStr)
	}
	chatSingleToReceiver.SenderInfo = getUserInfoServiceRes.Data

	if saveSingleChatMessageRes.IsOnline {

		zlog.Info("user[%v] to user[%v] on line chat message id[%v] data",
			singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, singleMessage.Message.MessageID)

		j.s.SendToClient(singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, msg_id, chatSingleToReceiver)
	} else {
		zlog.Info("user[%v] to user[%v] off line chat message id[%v] data",
			singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, chatSingleToReceiver.Data.MessageID)
	}

	return relay_api.NotifyUser(singleMessage.UserInfo.UserID, chatSingleToReceiver)
}

// 转发好友请求
func (j *Job) applyFriend(payload interface{}) error {
	applyFriendData := payload.(*event.ApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.ApplyFriendsToReceiver{})

	// 判断接收者是否在线
	getUserInfoReq := &api_account.GetUserInfoServiceReq{
		ClientInfo: applyFriendData.UserInfo,
		UserID:     applyFriendData.Message.ReceiverID,
	}
	getUserInfoRes, err := api_account.GetUserInfoService(context.Background(), getUserInfoReq)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but is err:%v", applyFriendData.UserInfo.UserID,
			applyFriendData.Message.ReceiverID, err)
		zlog.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Online {
		zlog.Info("user[%v] to on line user[%v] apply friend",
			applyFriendData.Message.SenderID, applyFriendData.Message.ReceiverID)

		return j.s.SendToClient(applyFriendData.Message.SenderID, applyFriendData.Message.ReceiverID, msg_id, applyFriendData.Message)
	}

	zlog.Info("user[%v] to off line user[%v] apply friend",
		applyFriendData.Message.SenderID, applyFriendData.Message.ReceiverID)

	return nil
}

// 转发同意好友申请
func (j *Job) agreeApplyFriend(payload interface{}) error {
	agreeApplyFriendData := payload.(*event.AgreeApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.AgreeApplyFriendsToReceiver{})

	// 判断接收者是否在线
	getUserInfoReq := &api_account.GetUserInfoServiceReq{
		ClientInfo: agreeApplyFriendData.UserInfo,
		UserID:     agreeApplyFriendData.Message.ReceiverID,
	}
	getUserInfoRes, err := api_account.GetUserInfoService(context.Background(), getUserInfoReq)
	if err != nil {
		errStr := fmt.Sprintf("user[%v] get user[%v] info, but is err:%v", agreeApplyFriendData.UserInfo.UserID,
			agreeApplyFriendData.Message.ReceiverID, err)
		zlog.Error(errStr)
		return fmt.Errorf(errStr)
	}

	if getUserInfoRes.Data.Status == im_home_proto.Enum_UserStatus_Enum_UserStatus_Online {
		zlog.Info("user[%v] to off line user[%v] agree apply friend",
			agreeApplyFriendData.Message.SenderID, agreeApplyFriendData.Message.ReceiverID)

		return j.s.SendToClient(agreeApplyFriendData.Message.SenderID, agreeApplyFriendData.Message.ReceiverID, msg_id, agreeApplyFriendData.Message)
	}

	zlog.Info("user[%v] to off line user[%v] agree apply friend",
		agreeApplyFriendData.Message.SenderID, agreeApplyFriendData.Message.ReceiverID)

	return nil
}
