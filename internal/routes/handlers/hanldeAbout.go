package handlers

import (
	"net/http"

	"github.com/Laelapa/GoHome/internal/interface/templates"
)

func (h *Handler) HandleGetAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// TODO: pull the parameters from env/secrets & request to pass them to the template instead of hardcoding them
	if err := templates.About("Laelapa - Demetrius Papas - About", "laelapa.fly.dev", "/about").Render(r.Context(), w); err != nil {
		h.LogError("Failed to render about page", r, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.LogInfo("Rendered: About page", r)
}
