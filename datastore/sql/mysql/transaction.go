package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/gobackerz/amagin/log"
)

type transaction struct {
	tx     *sql.Tx
	logger log.Logger
}

func (t *transaction) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.ExecContext(ctx, query, args...)
}

func (t *transaction) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.QueryContext(ctx, query, args...)
}

func (t *transaction) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *transaction) Commit() error {
	startTime := time.Now()

	defer t.queryLogger(startTime, "")

	return t.tx.Commit()
}

func (t *transaction) Rollback() error {
	startTime := time.Now()

	defer t.queryLogger(startTime, "")

	return t.tx.Rollback()
}

func (t *transaction) queryLogger(startTime time.Time, query string) {
	return
}
