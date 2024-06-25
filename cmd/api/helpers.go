package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"greenlight.flaviogalon.github.io/internal/validator"
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Limit the size of the request body to 1MB
	var maxBytes int64 = 1024 * 1024
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	// Set decoder to return error if JSON includes unknown fields
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf(
				"body contains badly-formed JSON (at characted %d)",
				syntaxError.Offset,
			)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf(
					"body contains incorrect JSON type for field %q",
					unmarshalTypeError.Field,
				)
			}
			return fmt.Errorf(
				"body contains incorrect JSON type (at characted %d)",
				unmarshalTypeError.Offset,
			)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Handling unknown JSON fields
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	// Calling Decode() again to make sure that there isn't any data left on Body
	err = dec.Decode(&struct{}{})
	// EOF is expected if JSON only has 1 value
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// Returns a string value from the query string, or the provided default value if no matching key could be found
func (app *application) readString(
	queryStringValues url.Values,
	key string,
	defaultValue string,
) string {
	s := queryStringValues.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

// Reads a string value from the query string and splits it into a slice on the comma.
// If no matching key could be found, returns the default value.
func (app *application) readCSV(
	queryStringValues url.Values,
	key string,
	defaultValue []string,
) []string {
	csv := queryStringValues.Get(key)
	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

// Reads a string value from the query string and converts it to an integer.
// If not matching key could be found returns the default value. If the conversion
// failed due to an error, the error will be added to the provided validator instance
func (app *application) readInt(
	queryStringValues url.Values,
	key string,
	defaultValue int,
	v *validator.Validator,
) int {
	s := queryStringValues.Get(key)
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}
