package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // never going to be serialized
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // serialized only if != 0
	Runtime   Runtime   `json:"runtime,omitempty"` // serialized only if != 0
	Genres    []string  `json:"genres,omitempty"`  // serialized only if != []
	Version   int32     `json:"version"`
}
