package main

import (
	"fmt"
	"net/http"
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

	fmt.Fprintf(w, "Should show the details of movie %d\n", id)
}
