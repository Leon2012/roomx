package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"roomx/recvsrv/pkg/consumer"
	"roomx/recvsrv/pkg/invoker"
	"roomx/recvsrv/pkg/server"
)

func main() {
	if err := ego.New().
		Invoker(
			invoker.Init,
			consumer.Init,
		).
		Serve(
			server.InitGRPC(),
			server.InitVernor(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
