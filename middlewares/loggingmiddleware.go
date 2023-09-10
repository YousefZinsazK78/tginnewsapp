package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	logg "github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeNow := time.Now()

		ctx.Next()

		timeEnd := time.Now()

		latencyTime := timeEnd.Sub(timeNow)

		reqMethod := ctx.Request.Method

		reqURI := ctx.Request.RequestURI

		statusCode := ctx.Writer.Status()

		clientIP := ctx.ClientIP()

		logg.WithFields(logg.Fields{
			"METHOD":    reqMethod,
			"URI":       reqURI,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("http request")

		ctx.Next()
	}
}
