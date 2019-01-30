package api

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/dreadatour/drone-station/pkg/dshttp"
	"github.com/dreadatour/drone-station/pkg/dslog"
)

// Config server
type Config struct {
	HTTP *dshttp.Config `envconfig:"HTTP"`
	Log  *dslog.Config  `envconfig:"LOG"`
}

// Run API server
func Run() error {
	// parse config
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		envconfig.Usage("", &cfg)
		return err
	}

	// initialise logger
	logger, err := dslog.NewLogger(cfg.Log)
	if err != nil {
		return err
	}

	logger.WithField("http", cfg.HTTP.Address()).Info("Hello, world!")

	return nil
}
