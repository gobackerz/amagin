package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Mysql struct {
	db      *sql.DB
	logger  Logger
	metrics Metrics
}

func New(config *Config, logger Logger, metrics Metrics) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error("Failed to open MySQL database: %v", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error("Failed to open MySQL database: %v", err)

		return nil, err
	}

	return &Mysql{db: db, logger: logger}, nil
}

func (m *Mysql) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.ExecContext(ctx, query, args...)
}

func (m *Mysql) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.QueryContext(ctx, query, args...)
}

func (m *Mysql) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	startTime := time.Now()

	defer m.queryLogger(startTime, query)

	return m.db.QueryRowContext(ctx, query, args...)
}

func (m *Mysql) Begin() (*Tx, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{tx: tx, logger: m.logger}, nil
}

func (m *Mysql) queryLogger(startTime time.Time, query string) {
	m.logger.Debug("query: %s", query)
}
