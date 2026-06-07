package httpprovisionerextension

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"
)

type provisionerExtension struct {
	logger        *zap.Logger
	endpoint      string
	poll_interval time.Duration

	// Calculate local md5
	local_config_path string
	local_MD5         string
}

func (e *provisionerExtension) Start(ctx context.Context, host component.Host) error {
	e.logger.Info("HTTP Provisioner Extension Started",
		zap.String("endpoint", e.endpoint),
		zap.Duration("poll_interval", e.poll_interval),
	)

	data, err := os.ReadFile(e.local_config_path)
	if err != nil {
		e.logger.Warn("Failed to read local config file" + e.local_config_path)
		return nil
	}

	e.local_MD5 = md5sum(data)
	e.logger.Info("Local md5sum: " + e.local_MD5)

	// Initial Poll
	e.poll()

	return nil
}

func (e *provisionerExtension) Shutdown(ctx context.Context) error {
	e.logger.Info("HTTP Provisioner Extension Stopped")
	return nil
}

// Poll the HTTP Provisioner
func (e *provisionerExtension) poll() {
	e.logger.Info("Initial poll from the HTTP Provisioner: ",
		zap.String("endpoint", e.endpoint),
	)

	// Download the config file from the HTTP Provisioner
	response, err := http.Get(e.endpoint)
	if err != nil {
		e.logger.Warn("Failed to download config file from HTTP Provisioner: " + err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		e.logger.Warn("Failed to download config file from HTTP Provisioner: " + response.Status)
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		e.logger.Warn("Failed to read response body from HTTP Provisioner: " + err.Error())
		return
	}

	md5sum := md5sum(data)
	e.logger.Info("Downloaded md5sum: " + md5sum)
	// Compare the md5sum with the local md5sum
	if md5sum != e.local_MD5 {
		e.logger.Info("Config file has changed, updating local config file and reloading collector")
	} else {
		e.logger.Info("Config file has not changed, no action needed")
	}
	// If they are different, update the local config file and reload the collector
	// If they are the same, do nothing
}

func md5sum(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

var _ extension.Extension = (*provisionerExtension)(nil)
