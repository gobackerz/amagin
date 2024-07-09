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
	"github.com/gobackerz/amagin/response"
)

type App struct {
	config.Config
	datastore.Datastore

	e      *gin.Engine
	Logger pkgLogger.Logger
}

type Handler func(ctx *Context) (interface{}, error)

func New(logger pkgLogger.Logger) *App {
	if logger == nil {
		logger = log.New(os.Stdout, "", -1)
	}

	cfg := config.New(logger)
	ds, _ := datastore.New(cfg, logger)
	e := gin.Default()

	return &App{Config: cfg, Datastore: ds, e: e, Logger: logger}
}

func (a *App) GET(relativePath string, handler Handler) {
	a.e.GET(relativePath, func(c *gin.Context) {
		a.processHandler(c, handler, http.StatusOK, http.StatusNotFound)
	})
}

func (a *App) POST(relativePath string, handler Handler) {
	a.e.POST(relativePath, func(c *gin.Context) {
		a.processHandler(c, handler, http.StatusCreated)
	})
}

func (a *App) PUT(relativePath string, handler Handler) {
	a.e.PUT(relativePath, func(c *gin.Context) {
		a.processHandler(c, handler, http.StatusOK)
	})
}

func (a *App) PATCH(relativePath string, handler Handler) {
	a.e.PATCH(relativePath, func(c *gin.Context) {
		a.processHandler(c, handler, http.StatusOK)
	})
}

func (a *App) DELETE(relativePath string, handler Handler) {
	a.e.DELETE(relativePath, func(c *gin.Context) {
		a.processHandler(c, handler, http.StatusNoContent)
	})
}

func (a *App) Run() {
	httpPort := a.Get("HTTP_PORT", "8000")

	a.e.Run(fmt.Sprintf(":%s", httpPort))
}

func (a *App) processHandler(c *gin.Context, handler Handler, defaultStatusCodes ...int) {
	ctx := NewContext(c, a)

	res, err := handler(ctx)
	statusCode := a.getStatusCode(res, err, defaultStatusCodes...)
	resp := a.getResponse(res, err)

	switch c.ContentType() {
	case "text/plain":
		c.String(statusCode, "%v", resp)
	case "application/xml":
		c.XML(statusCode, resp)
	default:
		c.JSON(statusCode, resp)
	}
}

func (a *App) getStatusCode(res any, err error, defaultStatusCodes ...int) int {
	defSuccessCode, defErrCode := a.getDefaultStatusCodes(defaultStatusCodes...)

	if err != nil {
		return a.getStatusCodeFromObj(err, defErrCode)
	}

	return a.getStatusCodeFromObj(res, defSuccessCode)
}

func (a *App) getDefaultStatusCodes(defaultStatusCodes ...int) (int, int) {
	var defaultSuccessCode, defaultErrCode int

	switch len(defaultStatusCodes) {
	case 0:
		defaultSuccessCode = http.StatusOK
		defaultErrCode = http.StatusInternalServerError
	case 1:
		defaultSuccessCode = defaultStatusCodes[0]
		defaultErrCode = http.StatusInternalServerError
	default:
		defaultSuccessCode = defaultStatusCodes[0]
		defaultErrCode = defaultStatusCodes[1]
	}

	return defaultSuccessCode, defaultErrCode
}

func (a *App) getStatusCodeFromObj(obj any, defaultStatusCode int) int {
	if respObj, ok := obj.(response.WithStatusCode); ok {
		return respObj.StatusCode()
	}

	return defaultStatusCode
}

func (a *App) getResponse(res any, err error) any {
	if res != nil && err != nil {
		return map[string]any{
			"data":  res,
			"error": err,
		}
	} else if err != nil {
		return err
	} else {
		return res
	}
}
