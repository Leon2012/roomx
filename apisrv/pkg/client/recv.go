package client

import (
	"roomx/proto/message"

	"github.com/gotomicro/ego/client/egrpc"
)

var (
	RecvConn   *egrpc.Component
	RecvClient message.RecvSrvClient
)

func InitRecvClient() error {
	RecvConn = egrpc.Load("service.recv").Build()
	RecvClient = message.NewRecvSrvClient(RecvConn.ClientConn)
	return nil
}
