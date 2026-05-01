package handlers

import (
	"log/slog"

	"backend/internal/auth"
	"backend/internal/store"
)

type Handler struct {
	Storage store.Storage
	Auth    auth.Auth
	Logger  *slog.Logger
}
