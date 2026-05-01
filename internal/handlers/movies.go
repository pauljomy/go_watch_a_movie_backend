package handlers

import "net/http"

func (h *Handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.Movies.GetAllMovies()
	if err != nil {
		h.errorJSON(w, err)
		return
	}

	if err = h.writeJSON(w, http.StatusOK, movies); err != nil {
		h.serverError(w, r, err)
	}
}
