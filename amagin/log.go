package amagin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/log"
)

type performanceLogger struct {
	logger amagin.Logger
	isTerm bool
	log    log.Log
}

func (l *performanceLogger) Write(p []byte) (int, error) {
	l.log.Message = string(p)

	l.logger.Info(string(p), l.log)

	return len(p), nil
}

func (l *performanceLogger) formatter(params gin.LogFormatterParams) string {
	l.log = log.Log{
		TimeStamp:        params.TimeStamp,
		StatusCode:       params.StatusCode,
		Latency:          params.Latency,
		ClientIP:         params.ClientIP,
		Method:           params.Method,
		Path:             params.Path,
		ErrorMessage:     params.ErrorMessage,
		BodySize:         params.BodySize,
		Keys:             params.Keys,
		IsPerformanceLog: true,
	}

	return defaultLogFormatter(params)
}

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string

	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}
