package handlers

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/interface/templates"
)

func (h *Handler) HandleUnderConstruction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := templates.UnderConstruction().Render(r.Context(), w); err != nil {
		h.LogError("Failed to render under construction page", r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.LogInfo("Rendered: Under Construction page", r)
}
