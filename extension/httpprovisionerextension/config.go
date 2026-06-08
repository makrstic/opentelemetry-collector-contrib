package httpprovisionerextension

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

const (
	DefaultEndpoint = "http://127.0.0.1:8080/api/config/default"
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
