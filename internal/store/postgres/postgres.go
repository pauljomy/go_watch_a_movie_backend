package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type PostgresMovieStore struct {
	db *pgxpool.Pool
}

func NewPostgresMovieStore(db *pgxpool.Pool) *PostgresMovieStore {
	return &PostgresMovieStore{db: db}
}

type PostgresAuthStore struct {
	db *pgxpool.Pool
}

func NewPostgressAuthStore(db *pgxpool.Pool) *PostgresAuthStore {
	return &PostgresAuthStore{db: db}
}
