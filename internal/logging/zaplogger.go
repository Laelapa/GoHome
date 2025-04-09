package logging

import (
	"go.uber.org/zap"
)

func NewLogger(env string) (*zap.SugaredLogger, error) {

	var logger *zap.Logger
	var err error

	if env == "prod" || env == "production" {
		logger, err = zap.NewProduction() // TODO: setup production output
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
