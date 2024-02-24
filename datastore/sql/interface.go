package sql

import (
	"context"
	"database/sql"
)

type SQL interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	Begin() error
	Commit() error
	Rollback() error
}
