package main

import (
	"backend/internals/repository"
	"backend/internals/repository/dbrepo"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	DSN    string
	DB     repository.DataBaseRepo
}

func main() {

	var addr string
	var dsn string

	flag.StringVar(&addr, "addr", ":4000", "backend server port")
	flag.StringVar(&dsn, "dsn", "host=localhost user=movies password=movies dbname=movies sslmode=disable timezone=UTC port=5433 connect_timeout=5", " Postgres database connection string")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
		DSN:    dsn,
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
