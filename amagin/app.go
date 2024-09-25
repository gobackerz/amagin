package amagin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/constants"
	pkgLogger "github.com/gobackerz/amagin/log"
)

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

type App struct {
	*datastore

	e      *gin.Engine
	logger amagin.Logger
	config *Config
}

type Handler func(ctx *Context) (interface{}, error)

func New() *App {
	e := gin.New()

	return &App{e: e}
}

func Default() *App {
	logger := pkgLogger.New()
	cfg := newConfig(logger)
	performanceLogWriter := &performanceLogger{logger: logger, isTerm: logger.IsTerm()}
	ds := &datastore{}

	logger.SetLevel(getLogLevelFromEnv())
	gin.ForceConsoleColor()

	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: performanceLogWriter.formatter, Output: performanceLogWriter}))

	return &App{config: cfg, datastore: ds, e: e, logger: logger}
}

func getLogLevelFromEnv() int {
	level := map[string]int{
		levelDebug: 0,
		levelInfo:  1,
		levelWarn:  2,
		levelError: 3,
	}

	logLevel, ok := level[os.Getenv(constants.EnvLogLevel)]
	if !ok {
		return 1
	}

	return logLevel
}

func (a *App) Config() *Config {
	return a.config
}

func (a *App) Logger() amagin.Logger {
	return a.logger
}

func (a *App) UseLogger(logger amagin.Logger) {
	a.logger = logger
}

func (a *App) UseSQL(sql amagin.SQL) {
	a.datastore.sql = sql
}

func (a *App) Run() error {
	httpPort := a.config.Get("HTTP_PORT", "8000")

	return a.e.Run(fmt.Sprintf(":%s", httpPort))
}
