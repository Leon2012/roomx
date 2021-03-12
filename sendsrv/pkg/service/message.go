package service

import (
	"context"
	"github.com/gotomicro/ego/server/egrpc"
	"roomx/libs/queue"
	"roomx/model"
	"roomx/model/mysql"
	"roomx/proto/common"
	"roomx/proto/message"
	"roomx/sendsrv/pkg/invoker"
	"time"
)

type MessageService struct {
	Server *egrpc.Component
}

func (s *MessageService) Send(context context.Context, req *message.MessageSendReq) (*message.MessageSendResp, error) {
	var err error
	var data *model.Message
	var body []byte
	data = &model.Message{
		Uid:      req.Uid,
		Rid:      req.Rid,
		Type:     1,
		Content:  req.Content,
		Extra:    req.Extra,
		Dateline: time.Now().Unix(),
	}
	err = mysql.MessageCreate(invoker.Db, data)
	if err != nil {
		return &message.MessageSendResp{Resp: &common.Resp{
			Code: 500,
			Msg:  err.Error(),
		}, Mid: 0}, nil
	}

	body, err = data.ToBytes()
	if err == nil {
		if err = invoker.Producer.Publish(&queue.Message{
			Topic:     "new.message.remind",
			Body:      body,
			Timestamp: 0,
		}); err != nil {
			return &message.MessageSendResp{Resp: &common.Resp{
				Code: 500,
				Msg:  "remind failed",
			}, Mid: 0}, nil
		}
	}

	return &message.MessageSendResp{Resp: &common.Resp{
		Code: 0,
		Msg:  "",
	}, Mid: data.Id}, nil
}
