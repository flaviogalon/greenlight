package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	query := `
        INSERT INTO movies (title, year, runtime, genres)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, version`

	jsonGenres, err := json.Marshal(movie.Genres)
	if err != nil {
		return err
	}

	args := []any{movie.Title, movie.Year, movie.Runtime, jsonGenres}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).
		Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Fetch a specific record from the movies table
func (m MovieModel) Get(id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	var movie Movie
	var genresJSONString string

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE id = $1`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		&genresJSONString,
		&movie.Version,
	)
	// DB query erros
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// JSON parsing error
	err = json.Unmarshal([]byte(genresJSONString), &movie.Genres)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

// Update a specific record in the movies table
func (m MovieModel) Update(movie *Movie) error {
	query := `
        UPDATE movies
        SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
        WHERE id = $5 AND version = $6
        RETURNING version`

	jsonGenres, err := json.Marshal(movie.Genres)
	if err != nil {
		return err
	}

	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		jsonGenres,
		movie.ID,
		movie.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = m.DB.QueryRowContext(ctx, query, args...).Scan(&movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

// Delete a specific record from the movies table
func (m MovieModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM movies
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
