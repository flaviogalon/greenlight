package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type envelope map[string]any

// Read an ID from a HTTP request's parameters
func (app *application) readIdFromRequestParams(r *http.Request) (int64, error) {
	idString := r.PathValue("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil
}

// Write JSON to the given ResponseWriter
func (app *application) writeJSON(
	w http.ResponseWriter,
	status int,
	data envelope,
	headers http.Header,
) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Just to make it easier to view in terminal
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
