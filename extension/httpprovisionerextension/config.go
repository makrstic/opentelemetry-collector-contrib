package httpprovisionerextension

import (
	"go.opentelemetry.io/collector/component"
	"time"
)

const (
	DefaultEndpoint = "http://127.0.0.1:8080"
	DefaultPollInterval = 45 * time.Minute
	DefaultConfigPath = "/etc/otel/config.yaml"
)

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
	PollInterval time.Duration `mapstructure:"poll_interval"`
	LocalConfigPath string `mapstructure:"local_config_path"`
}


func (c *Config) Validate() error {
	return nil
}

var _ component.Config = (*Config)(nil)
