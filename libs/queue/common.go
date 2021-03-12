package queue

import (
	"context"
	"errors"
	nsq "github.com/nsqio/go-nsq"
	"time"
)

type Message struct {
	Topic     string `json:"topic,omitempty"`
	Body      []byte `json:"body,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

var (
	ErrParseMessage     = errors.New("parse message failed")
	ErrMessageParameter = errors.New("error message parameter")
	ErrLookupdAddrs     = errors.New("lookupdaddrs is empty")
	ErrConcurrency      = errors.New("faile concurrency num")
)

type RedisOption struct {
	Network      string
	Addr         string
	Active       int
	Idle         int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type Operator interface {
	Operate(context.Context, *Message) (bool, error)
}

type NsqOption struct {
	Topic               string
	Channel             string
	Addr                string   //nsqd addr
	LookupdAddrs        []string //nsqlookupd addrs
	DialTimeout         int
	ReadTimeout         int
	WriteTimeout        int
	MaxInFlight         int
	QueueMsgSize        int
	LookupdPollInterval int
	Concurrency         int
}

func InitNsqProducer(c *NsqOption) (*nsq.Producer, error) {
	var err error
	config := nsq.NewConfig()
	if err = config.Set("dial_timeout", time.Duration(c.DialTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("read_timeout", time.Duration(c.ReadTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("write_timeout", time.Duration(c.WriteTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("max_in_flight", c.MaxInFlight); err != nil {
		return nil, err
	}
	producer, err := nsq.NewProducer(c.Addr, config)
	return producer, err
}

func InitNsqConsumer(topic, channel string, c *NsqOption) (*nsq.Consumer, error) {
	var err error
	config := nsq.NewConfig()
	if err = config.Set("dial_timeout", time.Duration(c.DialTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("read_timeout", time.Duration(c.ReadTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("write_timeout", time.Duration(c.WriteTimeout)*time.Millisecond); err != nil {
		return nil, err
	}
	if err = config.Set("max_in_flight", c.MaxInFlight); err != nil {
		return nil, err
	}
	if err = config.Set("lookupd_poll_interval", time.Duration(c.LookupdPollInterval)*time.Second); err != nil {
		return nil, err
	}
	consumer, err := nsq.NewConsumer(topic, channel, config)
	return consumer, err
}
