package handlers

import (
	"net/http"
)

func (h *Handler) LogInfo(msg string, r *http.Request) {
	h.Logger.Infow(
		msg,
		"method", r.Method,
		"path", r.URL.Path,
		"remote_addr", r.RemoteAddr,
	)
}

func (h *Handler) LogError(msg string, r *http.Request, err error) {
	h.Logger.Errorw(
		msg,
		"method", r.Method,
		"path", r.URL.Path,
		"remote_addr", r.RemoteAddr,
		"error", err,
	)
}