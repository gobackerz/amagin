package log

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
)

type Log struct {
	Level            string         `json:"level"`
	TimeStamp        time.Time      `json:"timeStamp"`
	StatusCode       int            `json:"statusCode,omitempty"`
	Latency          time.Duration  `json:"latency,omitempty"`
	ClientIP         string         `json:"clientIP,omitempty"`
	Method           string         `json:"method,omitempty"`
	Path             string         `json:"path,omitempty"`
	Message          string         `json:"message,omitempty"`
	ErrorMessage     string         `json:"errorMessage,omitempty"`
	BodySize         int            `json:"bodySize,omitempty"`
	Keys             map[string]any `json:"keys,omitempty"`
	IsPerformanceLog bool           `json:"-"`
}

type logger struct {
	logger     *log.Logger
	level      int
	isTerminal bool
}

func New(level int) *logger {
	l := logger{}

	l.logger = log.New(os.Stdout, "", 0)
	l.level = level
	l.isTerminal = isTerminal(l.logger.Writer())

	return &l
}

func (l *logger) Debug(msg string, args ...any) {
	l.log(levelDebug, msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.log(levelInfo, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log(levelWarn, msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.log(levelError, msg, args...)
}

func (l *logger) log(level int, msg string, args ...any) {
	if level < l.level {
		return
	}

	levelName := map[int]string{
		levelDebug: "DEBU",
		levelInfo:  "INFO",
		levelWarn:  "WARN",
		levelError: "ERRO",
	}

	levelColor := map[int]string{
		levelDebug: "\u001B[90m",
		levelInfo:  "\u001B[34m",
		levelWarn:  "\u001B[33m",
		levelError: "\u001B[31m",
	}

	var log Log

	out := fmt.Sprintf(msg, args...)
	log = Log{Level: levelName[level], Message: out, TimeStamp: time.Now()}

	performanceLog, ok := l.getPerformanceLog(args...)
	if ok {
		performanceLog.Level = levelName[levelInfo]

		if !l.isTerminal {
			performanceLog.Message = ""
		}

		log = *performanceLog
	}

	if !l.isTerminal {
		_ = json.NewEncoder(l.logger.Writer()).Encode(log)

		return
	}

	l.logger.Print(fmt.Sprintf("%s %s| %s | %s\u001B[0m ", log.TimeStamp.Format("2006/01/02 - 15:04:05"),
		levelColor[level], log.Level, log.Message))
}

func (l *logger) getPerformanceLog(args ...any) (*Log, bool) {
	if len(args) != 1 {
		return nil, false
	}

	log, ok := args[0].(Log)

	return &log, ok && log.IsPerformanceLog
}

func (l *logger) IsTerm() bool {
	return l.isTerminal
}

func isTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
