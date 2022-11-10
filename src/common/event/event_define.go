package event

import (
	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

var (
	Event_SingleMessage = &rabbitmq.EventStruct{"single.message", SingleMessage{}}
)
