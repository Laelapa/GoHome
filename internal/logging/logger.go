package logging

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

// func NewLogger(env string) (*Logger, error) {

// 	var logger *zap.Logger
// 	var err error

// 	if env == "prod" || env == "production" {
// 		logger, err = zap.NewProduction() // TODO: setup production output
// 	} else {
// 		logger, err = zap.NewDevelopment()
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Logger{logger}, nil
// }

func NewLogger(env string) (*Logger, error) {
	var config zap.Config

	switch env {
	case "dev", "development":
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "test", "testing":
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Development = true
		config.Sampling = nil
	case "prod", "production":
		config = zap.NewProductionConfig()
		config.InitialFields = map[string]interface{}{
			"service": "gohome: Laelapa.dev",
		}
	default:
		return nil, errors.New("cannot initialize logger due to invalid environment")
		// TODO: add default config
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
