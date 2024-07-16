package middleware

import (
	"fmt"
	"go-hex-arch/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log *logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()

		errs := strings.ReplaceAll(ctx.Errors.String(), "\n", ";")
		logMessage := fmt.Sprintf("HTTP: Method %s; Path: %s; Latency: %s; Status: %d; ClientIP: %s;\033[31m%s\033[0m",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			time.Since(t).String(),
			ctx.Writer.Status(),
			ctx.ClientIP(),
			errs,
		)

		switch {
		case ctx.Writer.Status() >= 400 && ctx.Writer.Status() < 500:
			log.Warn(logMessage)
		case ctx.Writer.Status() >= 500:
			log.Error(logMessage)
		default:
			log.Info(logMessage)
		}
	}
}
