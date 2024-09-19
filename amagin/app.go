package amagin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/config"
	"github.com/gobackerz/amagin/constants"
	"github.com/gobackerz/amagin/datastore"
	pkgLogger "github.com/gobackerz/amagin/log"
)

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

type App struct {
	amagin.Config
	amagin.Datastore

	e      *gin.Engine
	logger amagin.Logger
}

type Handler func(ctx *Context) (interface{}, error)

func New() *App {
	e := gin.New()

	return &App{e: e}
}

func Default() *App {
	logger := pkgLogger.New(getLogLevelFromEnv())
	cfg := config.New(logger)
	performanceLogWriter := &performanceLogger{logger: logger, isTerm: logger.IsTerm()}
	ds, _ := datastore.New(cfg, logger)

	gin.ForceConsoleColor()

	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: performanceLogWriter.formatter, Output: performanceLogWriter}))

	return &App{Config: cfg, Datastore: ds, e: e, logger: logger}
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

func (a *App) Logger() amagin.Logger {
	return a.logger
}

func (a *App) UseLogger(logger amagin.Logger) {
	a.logger = logger
}

func (a *App) UseConfig(cfg amagin.Config) {
	a.Config = cfg
}

func (a *App) UseDatastore(ds amagin.Datastore) {
	a.Datastore = ds
}

func (a *App) Run() {
	httpPort := a.Get("HTTP_PORT", "8000")

	a.e.Run(fmt.Sprintf(":%s", httpPort))
}
