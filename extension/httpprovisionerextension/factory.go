package httpprovisionerextension

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

var Type = component.MustNewType("httpprovisioner")

func createDefaultConfig() component.Config {
        return &Config {
                Endpoint: DefaultEndpoint,
                PollInterval: DefaultPollInterval,
		LocalConfigPath: DefaultConfigPath,
        }
}

func createExtension(
	ctx context.Context,
	set extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {

	c := cfg.(*Config)

	// Basic checks for users config
	if c.Endpoint == "" {
		set.Logger.Warn("Empty endpoint, using default: " + DefaultEndpoint)
		c.Endpoint = DefaultEndpoint
	}

	if c.PollInterval <= 0 {
		set.Logger.Warn("Wrong poll interval, using default: 45m ")
		c.PollInterval = DefaultPollInterval
	}

	if c.LocalConfigPath == "" {
		set.Logger.Warn("Empty local config path, using default: /etc/otel/config.yaml")
		c.LocalConfigPath = DefaultConfigPath
	}

	return &provisionerExtension{
		logger: set.Logger,
		endpoint: c.Endpoint,
		poll_interval: c.PollInterval,
		local_config_path: c.LocalConfigPath,
		
	}, nil
}

func NewFactory() extension.Factory {
	return extension.NewFactory(
		Type,
		createDefaultConfig,
		createExtension,
		component.StabilityLevelDevelopment,
	)
}

