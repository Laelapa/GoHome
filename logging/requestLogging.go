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

	// nil check mainly in case of misconfigured tests
	if r == nil {
		return []zap.Field{
			zap.String("error", "nil request"),
		}
	}

	return []zap.Field{
		zap.String(FieldRemoteAddr, filetLogValue(getClientIP(r))),
		zap.String(FieldMethod, r.Method),
		zap.String(FieldPath, filetLogValue(r.URL.Path)),
		zap.String(FieldReferer, filetLogValue(r.Referer())),
	}
}