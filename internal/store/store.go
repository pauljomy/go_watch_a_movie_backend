package store

import "backend/internal/models"

type MovieStorer interface {
	GetAllMovies() ([]*models.Movie, error)
}

type AuthStorer interface {
	GetUserByEmail(email string) (*models.User, error)
}

type Storage struct {
	Movies MovieStorer
	Auth   AuthStorer
}
