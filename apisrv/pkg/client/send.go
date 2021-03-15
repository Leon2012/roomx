package client

import (
	"roomx/proto/message"
	"github.com/gotomicro/ego/client/egrpc"
)

var (
	SendConn   *egrpc.Component
	SendClient message.SendSrvClient
)

func InitSendClient() error {
	SendConn = egrpc.Load("service.send").Build()
	SendClient = message.NewSendSrvClient(SendConn.ClientConn)
	return nil
}
