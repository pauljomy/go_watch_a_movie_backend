package store

import "backend/internal/models"

type MovieStorer interface {
	GetAllMovies() ([]*models.Movie, error)
}

type Storage struct {
	Movies MovieStorer
}
