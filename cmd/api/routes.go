package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("POST /v1/movies", app.createMovieHandler)
	router.HandleFunc("GET /v1/movies/{id}", app.getMovieHandler)
	router.HandleFunc("PUT /v1/movies/{id}", app.updateMovieHandler)
	// Match all other requests to a generic not found response
	router.HandleFunc("/", app.notFoundResponse)

	return router
}
