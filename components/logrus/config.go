package logrus

type Config struct {
	Level            string
	Path             string
	TimestampFormat  string
	ReportCaller     bool
	DisableColors    bool
	FullTimestamp    bool
	DisableTimestamp bool
}

func DefaultConfig() *Config {
	return &Config{
		Level:            "debug",
		Path:             "",
		TimestampFormat:  "",
		ReportCaller:     false,
		DisableColors:    false,
		FullTimestamp:    true,
		DisableTimestamp: false,
	}
}
