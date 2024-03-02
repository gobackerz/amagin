package config

type Config interface {
	Get(key string, defaultVal ...string) string
	Set(key string, value string) error
	Unset(key string) error
}
