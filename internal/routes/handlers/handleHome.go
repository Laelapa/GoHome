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
		h.LogError("Failed to render home page", r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return

	}

	h.LogInfo("Rendered: Home page", r)
}
