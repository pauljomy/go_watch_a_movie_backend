package main

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"backend/internal/handlers"
	"backend/internal/auth"
	"backend/internal/db"
	"backend/internal/store"
	"backend/internal/store/postgres"
)

func main() {
	var addr string
	var dsn string
	var jwtSecret string
	var jwtIssuer string
	var jwtAudience string
	var cookieDomain string

	flag.StringVar(&addr, "addr", ":4000", "backend server port")
	flag.StringVar(&dsn, "dsn", "host=localhost user=movies password=movies dbname=movies sslmode=disable timezone=UTC port=5433 connect_timeout=5", "Postgres connection string")
	flag.StringVar(&jwtSecret, "jwt-secret", "verysecretkey", "JWT secret key")
	flag.StringVar(&jwtIssuer, "jwt-issuer", "example.com", "JWT issuer")
	flag.StringVar(&jwtAudience, "jwt-audience", "example.com", "JWT audience")
	flag.StringVar(&cookieDomain, "cookie-domain", "localhost", "Cookie domain")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	conn, err := db.OpenDB(dsn, 25, 25, "15m")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("database connection pool established")
	defer conn.Close()

	h := &handlers.Handler{
		Storage: store.Storage{
			Movies: postgres.NewPostgresMovieStore(conn),
		},
		Auth: auth.Auth{
			Issuer:        jwtIssuer,
			Audience:      jwtAudience,
			Secret:        jwtSecret,
			TokenExpiry:   time.Minute * 15,
			RefreshExpiry: time.Hour * 24,
			CookieDomain:  cookieDomain,
			CookiePath:    "/",
			CookieName:    "__HOST-refresh-token",
		},
		Logger: logger,
	}

	app := &application{
		logger:  logger,
		config:  config{addr: addr},
		handler: h,
	}

	logger.Info("server started", "addr", addr)

	if err = app.run(app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
