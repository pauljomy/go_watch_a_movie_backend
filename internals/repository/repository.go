package repository

import "backend/internals/models"

type DataBaseRepo interface {
	GetAllMovies() ([]*models.Movie, error)
}
