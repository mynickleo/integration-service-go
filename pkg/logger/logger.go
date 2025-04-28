package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger
}

func ZapError(err error) zap.Field {
	return zap.Error(err)
}

func ZapField(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
