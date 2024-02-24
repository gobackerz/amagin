package app

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	*App
}

func NewContext(c *gin.Context, app *App) *Context {
	return &Context{Context: c, App: app}
}
