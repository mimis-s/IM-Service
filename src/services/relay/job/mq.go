package job

import (
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	send_to "github.com/mimis-s/IM-Service/src/services/gateway/service"
	"github.com/mimis-s/IM-Service/src/services/home/service/seralize"
	"github.com/mimis-s/IM-Service/src/services/relay/service"
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

type Job struct {
	s *service.Service
}

func InitMQ(s *service.Service) *Job {
	url := "amqp://dev:dev123@localhost:5672/"
	durable := false

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

	chatSingleToReceiver := im_home_proto.ChatSingleToReceiver{
		Data: singleMessage.Message,
	}
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.ChatSingleToReceiver{})

	send_to.SendToUser(singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, msg_id, chatSingleToReceiver)

	return nil
}

// 转发好友请求
func (j *Job) applyFriend(payload interface{}) error {
	applyFriendData := payload.(*event.ApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.ApplyFriendsToReceiver{})

	send_to.SendToUser(applyFriendData.Message.SenderID, applyFriendData.Message.ReceiverID, msg_id, applyFriendData.Message)
	return nil
}

// 转发同意好友申请
func (j *Job) agreeApplyFriend(payload interface{}) error {
	agreeApplyFriendData := payload.(*event.AgreeApplyFriend)
	msg_id := seralize.GetMsgIdByStruct(im_home_proto.AgreeApplyFriendsToReceiver{})

	send_to.SendToUser(agreeApplyFriendData.Message.SenderID, agreeApplyFriendData.Message.ReceiverID, msg_id, agreeApplyFriendData.Message)
	return nil
}
