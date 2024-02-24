package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gobackerz/amagin/constants"

	"github.com/gobackerz/amagin/log"
)

type mysql struct {
	db     *sql.DB
	tx     *sql.Tx
	logger log.Logger
}

func New(host, port, user, password, name string, logger log.Logger) (*mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := sql.Open(constants.MYSQL, dsn)
	if err != nil {
		logger.Printf("Failed to open MySQL database: %v", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Printf("Failed to open MySQL database: %v", err)

		return nil, err
	}

	return &mysql{db: db, logger: logger}, nil
}

func (m *mysql) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	m.logger.Printf("%s", query)

	return m.db.QueryRowContext(ctx, query, args...)
}

func (m *mysql) QueryRowContextTx(ctx context.Context, query string, args ...any) *sql.Row {
	m.logger.Printf("%s", query)

	if m.tx != nil {
		m.tx.QueryRowContext(ctx, query, args...)
	}

	return m.db.QueryRowContext(ctx, query, args...)
}

func (m *mysql) Begin() error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	m.tx = tx

	return nil
}

func (m *mysql) Commit() error {
	if m.tx != nil {
		return m.tx.Commit()
	}

	return nil
}

func (m *mysql) Rollback() error {
	if m.tx != nil {
		return m.tx.Rollback()
	}

	return nil
}
