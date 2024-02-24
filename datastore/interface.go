package datastore

import "github.com/gobackerz/amagin/datastore/sql"

type Datastore interface {
	Db() sql.SQL
}
