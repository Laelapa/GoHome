package handlers

import "net/http"

func HandleGetHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to GoHome"))
}
