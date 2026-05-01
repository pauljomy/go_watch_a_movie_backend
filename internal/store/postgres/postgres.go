package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type PostgresMovieStore struct {
	db *pgxpool.Pool
}

func NewPostgresMovieStore(db *pgxpool.Pool) *PostgresMovieStore {
	return &PostgresMovieStore{db: db}
}
