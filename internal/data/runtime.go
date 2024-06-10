package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRunTimeFormat = errors.New("invalid runtime format")

type Runtime int32

// Return the JSON-encoded value for a movie's runtime
// Example: "<runtime> minutes"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d minutes", r)
	// A valid JSON string must be wrapped in double quotes
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil
}

// Custom JSON parser for Runtime
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// Expecting a string in the format "<runtime> mins"
	// 1. Remove double quotes
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRunTimeFormat
	}

	// 2. Split the string to isolate the number
	parts := strings.Split(unquotedJSONValue, " ")

	// 3. Check parts of the string
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRunTimeFormat
	}

	// 4. Parse the number part of the string to an int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRunTimeFormat
	}

	// 5. Convert the int32 into the custom Runtime
	*r = Runtime(i)

	return nil
}
