package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// Return the JSON-encoded value for a movie's runtime
// Example: "<runtime> minutes"
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d minutes", r)
	// A valid JSON string must be wrapped in double quotes
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil
}
