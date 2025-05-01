package logging

import "go.uber.org/zap"

func (l *Logger) LogAppInfo(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func (l *Logger) LogAppWarn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func (l *Logger) LogAppError(msg string, err error, fields ...zap.Field) {
	newFields := append([]zap.Field{zap.Error(err)}, fields...)
	l.Error(msg, newFields...)
}
