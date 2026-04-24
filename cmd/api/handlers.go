package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "success",
		Message: "Welcome to the home page!",
		Version: "1.0.0",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

}

func (app *application) GetAllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DB.GetAllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
