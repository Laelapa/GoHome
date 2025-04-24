package logging

import (
	"net/http"
	"go.uber.org/zap"
)

func (l *Logger) LogRequestInfo(msg string, r *http.Request) {
	l.Info(msg, l.buildRequestFields(r)...)
}

func (l *Logger) LogRequestWarn(msg string, r *http.Request) {
	l.Warn(msg, l.buildRequestFields(r)...)
}

func (l *Logger) LogRequestError(msg string, r *http.Request, err error) {
	reqFields := l.buildRequestFields(r)
	reqFields = append(reqFields, zap.Error(err))

	l.Error(msg, reqFields...)
}


func (l *Logger) buildRequestFields( r *http.Request) []zap.Field {
	return []zap.Field{
		zap.String("remoteAddr", getClientIP(r)),
		zap.String("method", r.Method),
		zap.String("path", filetLogValue(r.URL.Path)),
		zap.String("referer", filetLogValue(r.Referer())),
	}
}