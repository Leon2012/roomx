package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"roomx/components/queue/producer"
	"roomx/libs/queue"
)

var (
	Logger   *elog.Component
	Db       *egorm.Component
	Producer queue.Producer
)

func Init() error {
	Logger = elog.Load("logger.default").Build()
	Db = egorm.Load("mysql.roomx").Build()
	Producer = producer.Load("queue.producer.nsq").Build()
	return nil
}
