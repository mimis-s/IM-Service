package job

import (
	"context"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	"github.com/mimis-s/IM-Service/src/common/im_log"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/IM-Service/src/services/message/api_message"
	"github.com/mimis-s/IM-Service/src/services/relay/service"
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service, configOptions *boot_config.ConfigOptions) *Job {
	url := configOptions.BootConfigFile.MQ.Url
	durable := configOptions.BootConfigFile.MQ.Durable

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
		im_log.Warn("user[%v] to [%v] chat save is err:%v", singleMessage.Message.SenderID,
			singleMessage.Message.ReceiverID, err)
		return err
	}
	if saveSingleChatMessageRes.IsOnline {
		chatSingleToReceiver := &im_home_proto.ChatSingleToReceiver{
			Data: singleMessage.Message,
		}
		msg_id := seralize.GetMsgIdByStruct(im_home_proto.ChatSingleToReceiver{})

		return j.s.SendToClient(singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, msg_id, chatSingleToReceiver)
	}

	return nil
}

// 转发好友请求
func (j *Job) applyFriend(payload interface{}) error {
	applyFriendData := payload.(*event.ApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.ApplyFriendsToReceiver{})
	return j.s.SendToClient(applyFriendData.Message.SenderID, applyFriendData.Message.ReceiverID, msg_id, applyFriendData.Message)
}

// 转发同意好友申请
func (j *Job) agreeApplyFriend(payload interface{}) error {
	agreeApplyFriendData := payload.(*event.AgreeApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.AgreeApplyFriendsToReceiver{})

	return j.s.SendToClient(agreeApplyFriendData.Message.SenderID, agreeApplyFriendData.Message.ReceiverID, msg_id, agreeApplyFriendData.Message)
}
