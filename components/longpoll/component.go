package longpoll

import (
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egin"
	"github.com/jcuga/golongpoll"
	"context"
)

const PackageName = "server.egin.longpoll"

type Component struct {
	gin       *egin.Component
	name      string
	config    *Config
	lpManager *golongpoll.LongpollManager
	logger    *elog.Component
}

func newComponent(name string, config *Config, logger *elog.Component, gin *egin.Component) *Component {
	ops := golongpoll.Options{
		LoggingEnabled:                 false,
		MaxLongpollTimeoutSeconds:      config.MaxTimeoutSeconds,
		MaxEventBufferSize:             config.MaxEventBufferSize,
		EventTimeToLiveSeconds:         config.EventTimeToLiveSeconds,
		DeleteEventAfterFirstRetrieval: config.DeleteEventAfterFirstRetrieval,
	}
	lpManager, err := golongpoll.StartLongpoll(ops)
	if err != nil {
		logger.Panic("init longpoll mangager failed")
	}
	return &Component{
		gin:       gin,
		lpManager: lpManager,
		logger:    logger,
		name:      name,
		config:    config,
	}
}

func (c *Component) Name() string {
	return c.name
}

func (c *Component) PackageName() string {
	return PackageName
}

func (c *Component) Init() error {
	return c.gin.Init()
}

func (c *Component) Start() error {
	return c.gin.Start()
}

func (c *Component) Stop() error {
	c.lpManager.Shutdown()
	return c.gin.Stop()
}

func (c *Component) GracefulStop(ctx context.Context) error {
	c.lpManager.Shutdown()
	return c.gin.GracefulStop(ctx)
}

func (c *Component) Info() *server.ServiceInfo {
	return c.gin.Info()
}

func (c *Component) Gin() *egin.Component {
	return c.gin
}

func (c *Component) LongpollManager() *golongpoll.LongpollManager {
	return c.lpManager
}