package invoker

import (
	"github.com/gotomicro/ego/core/elog"
)

var (
	Logger *elog.Component
)

func Init() error {
	Logger = elog.Load("logger.default").Build()
	return nil
}
