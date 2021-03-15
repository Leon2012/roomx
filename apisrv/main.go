package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"roomx/apisrv/pkg/client"
	"roomx/apisrv/pkg/invoker"
	"roomx/apisrv/pkg/server"
)

func main() {
	if err := ego.New().
		Invoker(
			invoker.Init,
			client.InitRecvClient,
			client.InitSendClient,
			).
		Serve(
			server.InitLongpoll(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
