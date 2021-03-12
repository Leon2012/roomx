package queue

type Producer interface {
	PublishAsync(req *Message)
	Publish(req *Message) error
	Start() error
	Stop() error
}
