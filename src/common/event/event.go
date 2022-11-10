package event

import (
	"fmt"

	"github.com/mimis-s/golang_tools/mq/rabbitmq"
)

const (
	exchangeName = "user.message.relay"
)

// 生产者
type mqProducersStruct struct {
	p *rabbitmq.Producers
}

var mqProducers *mqProducersStruct

func InitProducers(url string, durable bool) error {
	p, err := rabbitmq.InitProducers(url, exchangeName, durable)
	if err != nil {
		fmt.Printf("err:%v", err)
		return err
	}
	mqProducers = &mqProducersStruct{p}
	return nil
}

func Publish(r *rabbitmq.EventStruct, payload interface{}) error {
	return mqProducers.p.Publish(r.RoutingKey, payload)
}

// 消费者
func InitConsumers(url string, durable bool, cQueue []*rabbitmq.ConsumersQueue) {
	err := rabbitmq.RegisterConsumers(url, durable, exchangeName, cQueue)
	if err != nil {
		panic(err)
	}
}
