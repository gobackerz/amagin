package datastore

import (
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/constants"
	"github.com/gobackerz/amagin/datastore/sql"
	"github.com/gobackerz/amagin/datastore/sql/mysql"
)

type datastore struct {
	sql sql.SQL
}

func New(cfg amagin.Config, logger amagin.Logger) (*datastore, error) {
	ds := &datastore{}
	dbDialect := cfg.Get("DB_DIALECT")

	switch dbDialect {
	case constants.MYSQL, constants.POSTGRES:
		host := cfg.Get("DB_HOST")
		port := cfg.Get("DB_PORT")
		user := cfg.Get("DB_USER")
		password := cfg.Get("DB_PASSWORD")
		name := cfg.Get("DB_NAME")

		sql, err := connectSQLDB(host, port, user, password, name, dbDialect, logger)
		if err != nil {
			return nil, err
		}

		ds.sql = sql
	}

	return ds, nil
}

func connectSQLDB(host, port, user, password, name, dialect string, logger amagin.Logger) (sql.SQL, error) {
	var (
		sql sql.SQL
		err error
	)

	user = url.QueryEscape(user)
	password = url.QueryEscape(password)

	switch dialect {
	case constants.MYSQL:
		sql, err = mysql.New(host, port, user, password, name, logger)
	}

	return sql, err
}

func (ds *datastore) Db() sql.SQL {
	return ds.sql
}
