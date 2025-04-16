package handlers

import "net/http"

func (h *Handler) HandleGetHealth(w http.ResponseWriter, r *http.Request) {

	h.LogInfo("Rendered: Health check", r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
