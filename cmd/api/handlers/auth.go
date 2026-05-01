package handlers

import (
	"errors"
	"net/http"

	"backend/internal/auth"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := h.readJSON(w, r, &requestPayload); err != nil {
		h.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	user, err := h.Storage.Auth.GetUserByEmail(requestPayload.Email)
	if err != nil {
		h.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {

		h.errorJSON(w, errors.New("Invalid Credentials"), http.StatusBadRequest)
		return
	}

	u := auth.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := h.Auth.GenerateTokenPair(&u)
	if err != nil {
		h.errorJSON(w, err)
		return
	}

	refreshCookie := h.Auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	h.writeJSON(w, http.StatusAccepted, tokens)
}
