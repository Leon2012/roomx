package server

import (
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egrpc"
	"roomx/proto/message"
	"roomx/sendsrv/pkg/service"
)

var (
	GRPC *egrpc.Component
)

func InitGRPC() server.Server {
	GRPC = egrpc.Load("server.grpc").Build()
	message.RegisterSendSrvServer(GRPC.Server, &service.MessageService{Server: GRPC})
	return GRPC
}
