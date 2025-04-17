package handlers

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/interface/templates"
)

func (h *Handler) HandleGetStack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := templates.Stack("Laelapa - Demetrius Papas - Tech Stack", "laelapa.fly.dev", "/stack").Render(r.Context(), w); err != nil {
		h.LogError("Failed to render stack page", r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.LogInfo("Rendered: Stack page", r)
}