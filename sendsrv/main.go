package main

import (
	"roomx/components/logrus"
	"roomx/sendsrv/pkg/invoker"
	"roomx/sendsrv/pkg/server"

	"github.com/gotomicro/ego"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Serve(
			server.InitGRPC(),
			server.InitVernor(),
		).
		Run(); err != nil {
		logrus.Fatal(err.Error())
	}
}
