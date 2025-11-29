package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
}

func (s Sqlite) Connect(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "sqlc-tutorial.sqlite")
	if err != nil {
		return nil, err
	}

	return db, nil
}
