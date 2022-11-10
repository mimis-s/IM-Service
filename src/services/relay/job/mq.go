package job

import (
	"github.com/mimis-s/IM-Service/src/common/commonproto/im_home_proto"
	"github.com/mimis-s/IM-Service/src/common/event"
	send_to "github.com/mimis-s/IM-Service/src/services/gateway/service"
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

	send_to.SendToUser(singleMessage.Message.SenderID, singleMessage.Message.ReceiverID, chatSingleToReceiver)

	return nil
}
