package longpoll

import (
	"github.com/jcuga/golongpoll"
)

type Config struct {
	MaxTimeoutSeconds              int
	MaxEventBufferSize             int
	EventTimeToLiveSeconds         int
	DeleteEventAfterFirstRetrieval bool
}

func DefaultConfig() *Config {
	return &Config{
		MaxTimeoutSeconds:              120,
		MaxEventBufferSize:             250,
		EventTimeToLiveSeconds:         golongpoll.FOREVER,
		DeleteEventAfterFirstRetrieval: false,
	}
}

