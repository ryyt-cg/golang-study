package db

import (
	"context"
	"database/sql"
)

type DBase interface {
	Connect(context.Context) (*sql.DB, error)
}
