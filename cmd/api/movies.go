package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.flaviogalon.github.io/internal/data"
)

// Create a new Movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new Movie")
}

// Get a Movie by ID
func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdFromRequestParams(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
