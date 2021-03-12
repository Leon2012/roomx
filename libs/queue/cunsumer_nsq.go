package queue

import "github.com/nsqio/go-nsq"

type NsqConsumer struct {
	topic, channel string
	consumer       *nsq.Consumer
	o              *NsqOption
	h              HandlerGenerator
}

func NewNsqConsumer(topic, channel string, h HandlerGenerator, o *NsqOption) (*NsqConsumer, error) {
	consumer, err := InitNsqConsumer(topic, channel, o)
	if err != nil {
		return nil, err
	}
	c := &NsqConsumer{
		topic:    topic,
		channel:  channel,
		consumer: consumer,
		o:        o,
		h:        h,
	}
	consumer.AddHandler(c)
	err = consumer.ConnectToNSQLookupds(o.LookupdAddrs)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (n *NsqConsumer) HandleMessage(message *nsq.Message) error {
	if n.h == nil {
		return nil
	}
	payload := &Payload{n.topic, n.channel, message}
	handler, err := n.h(payload)
	if err != nil {
		return nil
	}
	if err = handler.Perform(); err != nil {
		return err
	}
	return nil
}

func (n *NsqConsumer) GetConsumer() *nsq.Consumer  {
	return n.consumer
}