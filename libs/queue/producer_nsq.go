package queue

import (
	"errors"
	"github.com/micro/go-log"
	"github.com/nsqio/go-nsq"
	//"encoding/json"
	"roomx/libs/waitgroup"
)

type NsqProducer struct {
	waitgroup waitgroup.Wrapper
	msgChan   chan *Message
	exitChan  chan int
	producer  *nsq.Producer
	o         *NsqOption
}

func NewNsqProducer(o *NsqOption) (*NsqProducer, error) {
	producer, err := InitNsqProducer(o)
	if err != nil {
		return nil, err
	}
	return &NsqProducer{
		producer: producer,
		msgChan:  make(chan *Message, o.QueueMsgSize),
		exitChan: make(chan int),
	}, nil
}

func (p *NsqProducer) messagePump() {
	var (
		msg *Message
		//body []byte
		err error
		//json = jsoniter.ConfigCompatibleWithStandardLibrary
	)
	for {
		select {
		case msg = <-p.msgChan:
			//body, err = json.Marshal(msg)
			//if err == nil {
			err = p.producer.Publish(msg.Topic, msg.Body)
			//}
			if err != nil {
				log.Fatalf("producer publish message failed, error : %s", err.Error())
			}
		case <-p.exitChan:
			goto exit
		}
	}
exit:
	log.Log("producer message pump exist")
}

func (p *NsqProducer) Start() error {
	p.waitgroup.Wrap(func() {
		p.messagePump()
	})
	return nil
}

func (p *NsqProducer) Publish(msg *Message) error {
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//body, err := json.Marshal(msg)
	//if err != nil {
	//	return err
	//}
	if msg.Topic == "" || len(msg.Body) == 0 {
		return errors.New("empty data")
	}
	return p.producer.Publish(msg.Topic, msg.Body)
}

func (p *NsqProducer) PublishAsync(msg *Message) {
	if msg.Topic == "" || len(msg.Body) == 0 {
		return
	}
	select {
	case p.msgChan <- msg:
	default:

	}
}

func (p *NsqProducer) Stop() error {
	p.producer.Stop()
	close(p.exitChan)
	// synchronize the close of messagePump()
	p.waitgroup.Wait()
	return nil
}
