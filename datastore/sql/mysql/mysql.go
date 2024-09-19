package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/constants"
)

type mysql struct {
	db     *sql.DB
	logger amagin.Logger
}

func New(host, port, user, password, name string, logger amagin.Logger) (*mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := sql.Open(constants.MYSQL, dsn)
	if err != nil {
		logger.Error("Failed to open MySQL database: %v", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error("Failed to open MySQL database: %v", err)

		return nil, err
	}

	return &mysql{db: db, logger: logger}, nil
}

func (m *mysql) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.ExecContext(ctx, query, args...)
}

func (m *mysql) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.QueryContext(ctx, query, args...)
}

func (m *mysql) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.QueryRowContext(ctx, query, args...)
}

func (m *mysql) Begin() (pkgSQL.Transaction, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	return &transaction{tx: tx, logger: m.logger}, nil
}

func (m *mysql) queryLogger(startTime time.Time, query string) {
	m.logger.Debug("%s", query)
}
