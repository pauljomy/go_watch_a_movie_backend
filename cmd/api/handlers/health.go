package handlers

import "net/http"

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	h.writeJSON(w, http.StatusOK, payload)
}
