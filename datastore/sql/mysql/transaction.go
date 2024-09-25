package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Tx struct {
	tx      *sql.Tx
	logger  Logger
	metrics Metrics
}

func Begin(db any) (*Tx, error) {
	sqlDb, ok := db.(*Mysql)
	if !ok {
		return nil, fmt.Errorf("unexpected instance")
	}

	tx, err := sqlDb.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx, logger: sqlDb.logger, metrics: sqlDb.metrics}, nil
}

func (t *Tx) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.ExecContext(ctx, query, args...)
}

func (t *Tx) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.QueryContext(ctx, query, args...)
}

func (t *Tx) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	startTime := time.Now()

	defer t.queryLogger(startTime, query)

	return t.tx.QueryRowContext(ctx, query, args...)
}

func (t *Tx) Commit() error {
	startTime := time.Now()

	defer t.queryLogger(startTime, "")

	return t.tx.Commit()
}

func (t *Tx) Rollback() error {
	startTime := time.Now()

	defer t.queryLogger(startTime, "")

	return t.tx.Rollback()
}

func (t *Tx) queryLogger(startTime time.Time, query string) {
	return
}
