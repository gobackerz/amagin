package amagin

import (
	"github.com/gin-gonic/gin"

	"github.com/gobackerz/amagin"
)

type Context struct {
	*gin.Context
	*datastore

	logger amagin.Logger
}

func NewContext(c *gin.Context, ds *datastore, logger amagin.Logger) *Context {
	return &Context{Context: c, datastore: ds, logger: logger}
}

func (c *Context) Logger() amagin.Logger {
	return c.logger
}
