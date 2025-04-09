package handlers

import "net/http"

func (h *Handler) HandleGetHealth(w http.ResponseWriter, r *http.Request) {

	h.Logger.Infow("Rendered: Health check",
		"method", r.Method,
		"path", r.URL.Path,
		"remote_addr", r.RemoteAddr,
	)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
