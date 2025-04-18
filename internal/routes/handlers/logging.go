package handlers

import (
	"net/http"
	"strings"
)

func (h *Handler) LogInfo(msg string, r *http.Request) {
	h.Logger.Infow(
		msg,
		"method", r.Method,
		"path", sanitizeLogValue(r.URL.Path),
		"remote_addr", getClientIP(r),
		"referer", sanitizeLogValue(r.Referer()),
	)
}

func (h *Handler) LogError(msg string, r *http.Request, err error) {
	h.Logger.Errorw(
		msg,
		"method", r.Method,
		"path", sanitizeLogValue(r.URL.Path),
		"remote_addr", getClientIP(r),
		"referer", sanitizeLogValue(r.Referer()),
		"error", err,
	)
}

// TODO: Consider generalizing by using X-forwarded-for instead of Fly.io specific header.
//
// getClientIP retrieves the client IP address from the request if a reverse proxy is sitting in the middle.
// Currently only works if deployed on fly.io.
// SECURITY: Fly-Client-IP can obviously be spoofed in non-fly.io environments.
func getClientIP(r *http.Request) string {
	// Check if the request has a fly.io forwarded header
	if clientIP := r.Header.Get("Fly-Client-IP"); clientIP != "" {
		return clientIP
	}
	// Fallback to the remote address
	return "fly.io r-proxy: " + r.RemoteAddr
}

func sanitizeLogValue(v string) string {
	replacer := strings.NewReplacer(
		"\n", " [newline]",
		"\r", " [carriage return]",
		"\t", " [tab]",
		"<", " [less than]",
		">", " [greater than]",
		"\u001b", " [escape]",
	)
	return replacer.Replace(v)
}
