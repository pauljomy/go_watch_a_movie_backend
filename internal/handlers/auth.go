package handlers

import (
	"net/http"

	"backend/internal/auth"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := h.readJSON(w, r, &requestPayload); err != nil {
		h.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	u := auth.JwtUser{
		ID:        1,
		FirstName: "Paul",
		LastName:  "Jomy",
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
