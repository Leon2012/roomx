package consumer

import (
	"roomx/components/queue/consumer"
	"roomx/libs/queue"
)

var (
	handle   *Handler = &Handler{}
	Consumer *consumer.Component
)

func Init() error {
	Consumer = consumer.Load("queue.consumer.nsq").Build(func(payload *queue.Payload) (queue.Handler, error) {
		handle.SetPayload(payload)
		return handle, nil
	})
	//disable log
	Consumer.GetConsumer().SetLogger(nil, 0)
	return nil
}
