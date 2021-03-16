package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
	"roomx/components/logrus"
)

var (
	Logger *logrus.Component
	Db     *egorm.Component
	Redis  *eredis.Component
)

func Init() error {
	Logger = logrus.Load("logger.logrus").Build()
	Db = egorm.Load("mysql.roomx").Build()
	Redis = eredis.Load("redis.roomx").Build()
	return nil
}
