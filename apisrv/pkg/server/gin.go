package server

import (
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egin"
)

var (
	Gin *egin.Component
)

func InitGin() server.Server {
	Gin = egin.Load("server.http").Build()
	return Gin
}