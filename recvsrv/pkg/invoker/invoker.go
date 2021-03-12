package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
	"github.com/gotomicro/ego/core/elog"
)

var (
	Logger *elog.Component
	Db     *egorm.Component
	Redis  *eredis.Component
)

func Init() error {
	Logger = elog.Load("logger.default").Build()
	Logger.IsDebugMode()
	Db = egorm.Load("mysql.roomx").Build()
	Redis = eredis.Load("redis.roomx").Build()
	return nil
}
