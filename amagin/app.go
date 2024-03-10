package amagin

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gobackerz/amagin/config"
	"github.com/gobackerz/amagin/datastore"
	pkgLogger "github.com/gobackerz/amagin/log"
)

type App struct {
	config.Config
	datastore.Datastore

	e      *gin.Engine
	Logger pkgLogger.Logger
}

func New(logger pkgLogger.Logger) *App {
	if logger == nil {
		logger = log.New(os.Stdout, "", -1)
	}

	cfg := config.New(logger)
	ds, _ := datastore.New(cfg, logger)
	e := gin.Default()

	return &App{Config: cfg, Datastore: ds, e: e, Logger: logger}
}

func (a *App) GET(relativePath string, handler func(ctx *Context) (interface{}, error)) {
	a.e.GET(relativePath, func(c *gin.Context) {
		ctx := NewContext(c, a)

		res, err := handler(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func (a *App) POST(relativePath string, handler func(ctx *Context) (interface{}, error)) {
	a.e.GET(relativePath, func(c *gin.Context) {
		ctx := NewContext(c, a)

		res, err := handler(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func (a *App) PUT(relativePath string, handler func(ctx *Context) (interface{}, error)) {
	a.e.GET(relativePath, func(c *gin.Context) {
		ctx := NewContext(c, a)

		res, err := handler(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func (a *App) PATCH(relativePath string, handler func(ctx *Context) (interface{}, error)) {
	a.e.GET(relativePath, func(c *gin.Context) {
		ctx := NewContext(c, a)

		res, err := handler(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func (a *App) DELETE(relativePath string, handler func(ctx *Context) (interface{}, error)) {
	a.e.GET(relativePath, func(c *gin.Context) {
		ctx := NewContext(c, a)

		res, err := handler(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func (a *App) Run() {
	httpPort := a.Get("HTTP_PORT", "8000")

	a.e.Run(fmt.Sprintf(":%s", httpPort))
}
