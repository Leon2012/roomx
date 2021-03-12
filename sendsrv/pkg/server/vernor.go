package server

import (
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egovernor"
)

var (
	Vernor *egovernor.Component
)

func InitVernor() server.Server {
	Vernor = egovernor.Load("server.governor").Build()
	return Vernor
}
