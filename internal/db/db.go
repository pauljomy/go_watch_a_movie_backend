package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDB(dsn string, maxIdleConns, maxOpenConns int, maxIdleTime string) (*pgxpool.Pool, error) {

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(maxOpenConns)
	config.MinConns = int32(maxIdleConns)
	config.MaxConnIdleTime = duration

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
