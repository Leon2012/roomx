package queue

import "github.com/nsqio/go-nsq"

type HandlerGenerator func(*Payload) (Handler, error)

type Handler interface {
	Id() string
	Perform() error
}

type Payload struct {
	Topic   string
	Channel string
	Message *nsq.Message
}
