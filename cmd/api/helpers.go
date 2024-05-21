package main

import (
	"errors"
	"net/http"
	"strconv"
)

// Read an ID from a HTTP request's parameters
func (app *application) readIdFromRequestParams(r *http.Request) (int64, error) {
	idString := r.PathValue("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil
}
