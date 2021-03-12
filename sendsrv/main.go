package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"roomx/sendsrv/pkg/invoker"
	"roomx/sendsrv/pkg/server"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Serve(
			server.InitGRPC(),
			server.InitVernor(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
