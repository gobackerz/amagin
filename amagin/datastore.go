package amagin

import (
	"database/sql"
	"fmt"

	"github.com/gobackerz/amagin"
)

type datastore struct {
	sql amagin.SQL
}

func newDatastore(cfg *Config, logger amagin.Logger) (*datastore, error) {
	ds := &datastore{}
	sqliteDB := cfg.Get("SQLITE_DB")

	if sqliteDB != "" {
		logger.Info("connecting to sqlite: ", sqliteDB)

		db, err := sql.Open("sqlite", sqliteDB)
		if err != nil {
			return nil, fmt.Errorf("failed to open sqlite connection: %v", err)
		}

		logger.Info("connected to sqlite: ", sqliteDB)

		ds.sql = db
	}

	return ds, nil
}

func (ds *datastore) SQL() amagin.SQL {
	return ds.sql
}
