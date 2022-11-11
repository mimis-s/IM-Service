package event

import (
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

var (
	Event_SingleMessage    = &rabbitmq.EventStruct{"user.message.single", SingleMessage{}}
	Event_ApplyFriend      = &rabbitmq.EventStruct{"user.message.apply.friend", ApplyFriend{}}
	Event_AgreeApplyFriend = &rabbitmq.EventStruct{"user.message.agree.apply.friend", AgreeApplyFriend{}}
)
