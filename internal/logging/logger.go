package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(env string) (*Logger, error) {
	var config zap.Config

	switch env {
	case "dev", "development":
		config = setupDevConfig()
	case "test", "testing":
		config = setupTestConfig()
	case "prod", "production":
		config = setupProdConfig()
	default:
		env = "dev"
		fmt.Fprintf(os.Stderr, "WARNING: Unknown environment, defaulting to development\n")
		config = setupDevConfig()
	}

	logger, err := config.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	logger.Info(
		"Logger initialized",
		zap.String("environment", env),
		zap.String("level", config.Level.String()),
	)

	return &Logger{logger}, nil
}

func setupDevConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config
}

func setupTestConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.Development = true
	config.Sampling = nil

	return config
}

func setupProdConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.InitialFields = map[string]interface{}{
		"service": "gohome: Laelapa.dev",
	}
	return config
}
