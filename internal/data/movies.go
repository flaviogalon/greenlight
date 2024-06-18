package data

import (
	"database/sql"
	"time"

	"greenlight.flaviogalon.github.io/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // never going to be serialized
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // serialized only if != 0
	Runtime   Runtime   `json:"runtime,omitempty"` // serialized only if != 0
	Genres    []string  `json:"genres,omitempty"`  // serialized only if != []
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	// Title must not be empty
	v.Check(movie.Title != "", "title", "must be provided")
	// Title most be at most 500 bytes long
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	// Year must be provided
	v.Check(movie.Year != 0, "year", "must be provided")
	// Year must be greater than 1888
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	// Year must not be in the future
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "most not be in the future")

	// Runtime must be provided
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	// Runtime must be positive
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	// Genres must be provided
	v.Check(movie.Genres != nil, "genres", "must be provided")
	// Genres must contain at least 1 element
	v.Check(len(movie.Genres) > 0, "genres", "must container at least 1 genre")
	// Genres must containt at most 5 elemenets
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	// Genres must contain unique elements
	v.Check(validator.Unique(movie.Genres), "genres", "most not contain duplicate values")
}

type MovieModel struct {
	DB *sql.DB
}

// Insert a new record in the movies table
func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

// Fetch a specific record from the movies table
func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Update a specific record in the movies table
func (m MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete a specific record from the movies table
func (m MovieModel) Delete(id int64) error {
	return nil
}
