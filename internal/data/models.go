package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
}

type ModelsConfig struct {
	DBQueryTimeout time.Duration
}

func NewModels(db *sql.DB, timeout time.Duration) Models {
	modelsConfig := ModelsConfig{DBQueryTimeout: timeout}
	return Models{
		Movies: MovieModel{DB: db, ModelsConfig: modelsConfig},
	}
}
