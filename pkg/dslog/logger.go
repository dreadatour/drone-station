package dslog

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// NewLogger creates new logger
func NewLogger(cfg *Config) (*logrus.Logger, error) {
	switch strings.ToLower(cfg.Level) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		return nil, fmt.Errorf("Error parse config: unknown log level '%s'", cfg.Level)
	}

	switch strings.ToLower(cfg.Format) {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	case "plaintext":
		logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		return nil, fmt.Errorf("Error parse config: unknown log format '%s'", cfg.Format)
	}

	logger := logrus.StandardLogger()

	return logger, nil
}
