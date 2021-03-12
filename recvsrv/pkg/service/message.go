package service

import (
	"context"
	"github.com/gotomicro/ego/server/egrpc"
	"roomx/model"
	"roomx/model/mysql"
	"roomx/model/redis"
	"roomx/proto/message"
	"roomx/recvsrv/pkg/invoker"
)

type MessageService struct {
	Server *egrpc.Component
}

func (s *MessageService) Recv(context context.Context, req *message.MessageRecvReq) (*message.MessageRecvResp, error) {
	var (
		xMessage       *model.Message
		messages       model.Messages
		currId, nextId int32
		err            error
		resp           *message.MessageRecvResp = &message.MessageRecvResp{}
		xModel          *message.Model
	)
	currId = req.Seq
	if currId > 0 {
		invoker.Logger.Infof("message seq : %d", req.Seq)
		xMessage, err = mysql.MessageGet(invoker.Db, currId)
		if err != nil {
			invoker.Logger.Error("call MessageGet failed, error : " + err.Error())
			return nil, err
		}
		invoker.Logger.Infof("get message : %s", xMessage.Content)
	}
	messages, nextId, err = redis.MessagesNext(invoker.Redis, req.Uid, req.Rid, currId, xMessage)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(messages); i++ {
		xMessage = messages[i]
		xModel = &message.Model{
			Id:       xMessage.Id,
			Uid:      xMessage.Uid,
			Rid:      xMessage.Rid,
			Type:     xMessage.Type,
			Content:  xMessage.Content,
			Extra:    xMessage.Extra,
			Dateline: xMessage.Dateline,
		}
		resp.Messages = append(resp.Messages, xModel)
	}
	resp.Nextseq = nextId
	return resp, nil
}
