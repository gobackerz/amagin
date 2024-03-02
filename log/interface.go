package log

type Logger interface {
	Fatalf(format string, v ...any)
	Printf(format string, v ...any)
}
