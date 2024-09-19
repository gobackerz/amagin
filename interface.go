package amagin

import "github.com/gobackerz/amagin/datastore/sql"

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type Config interface {
	Get(key string, defaultVal ...string) string
	Set(key string, value string) error
	Unset(key string) error
}

type Datastore interface {
	Db() sql.SQL
}
