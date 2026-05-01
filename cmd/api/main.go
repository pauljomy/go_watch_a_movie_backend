package main

import (
	"backend/internals/repository"
	"backend/internals/repository/dbrepo"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *slog.Logger
	DSN    string
	DB     repository.DataBaseRepo
	auth   Auth
	Domain string
}

func main() {

	var addr string
	var dsn string
	var jwtSecret string
	var jwtIssuer string
	var jwtAudience string
	var cookieDomain string
	var domain string

	flag.StringVar(&addr, "addr", ":4000", "backend server port")
	flag.StringVar(&dsn, "dsn", "host=localhost user=movies password=movies dbname=movies sslmode=disable timezone=UTC port=5433 connect_timeout=5", " Postgres database connection string")

	flag.StringVar(&jwtSecret, "jwt-secret", "verysecretkey", "JWT secret key")
	flag.StringVar(&jwtIssuer, "jwt-issuer", "example.com", "JWT issuer")
	flag.StringVar(&jwtAudience, "jwt-audience", "example.com", "JWT audience")
	flag.StringVar(&cookieDomain, "cookie-domain", "localhost", "Cookie domain")
	flag.StringVar(&domain, "domain", "example.com", "Domain")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	auth := Auth{
		Issuer:        jwtIssuer,
		Audience:      jwtAudience,
		Secret:        jwtSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  cookieDomain,
		CookiePath:    "/",
		CookieName:    "__HOST-refresh-token",
	}

	app := &application{
		logger: logger,
		DSN:    dsn,
		auth:   auth,
		Domain: domain,
	}

	conn, err := app.connectToDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	defer conn.Close()

	logger.Info("server started at port:", "addr", addr)

	err = http.ListenAndServe(addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
