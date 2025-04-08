package handlers

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/interface/templates"
)

func (h *Handler) HandleGetHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := templates.Home().Render(r.Context(), w); err != nil {
		h.Logger.Errorw("Failed to render home page",
		"method", r.Method,
		"path", r.URL.Path,
		"remote_addr", r.RemoteAddr,
		"error", err,
	)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	h.Logger.Infow("Rendered: Home page",
	"method", r.Method,
	"path", r.URL.Path,
	"remote_addr", r.RemoteAddr,
	)
}
