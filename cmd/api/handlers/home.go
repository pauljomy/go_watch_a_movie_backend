package handlers

import "net/http"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "success",
		Message: "Welcome to the home page!",
		Version: "1.0.0",
	}

	if err := h.writeJSON(w, http.StatusOK, payload); err != nil {
		h.serverError(w, r, err)
	}
}
