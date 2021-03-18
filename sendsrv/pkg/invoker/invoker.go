package invoker

import (
	"roomx/components/logrus"
	"roomx/components/queue/producer"
	"roomx/libs/queue"

	"github.com/gotomicro/ego-component/egorm"
)

var (
	Logger   *logrus.Component
	Db       *egorm.Component
	Producer queue.Producer
)

func Init() error {
	Logger = logrus.Load("logger.logrus").Build()
	Db = egorm.Load("mysql.roomx").Build()
	Producer = producer.Load("queue.producer.nsq").Build()
	return nil
}
