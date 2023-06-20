package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
)

// RequestIDHeaderValue value for the request id header
// const RequestIDHeaderValue = "X-Request-ID"

// LoggerMiddleware middleware for logging incoming requests
func LoggerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
			start := time.Now()

			logger, err := hertzZerolog.GetLogger()
			if err != nil {
					hlog.Error(err)
					ctx.Next(c)
					return
			}
			
			// reqId := c.Value(RequestIDHeaderValue).(string)
			// if reqId != "" {
			// 		logger = logger.WithField("request_id", reqId)
			// }
			
			c = logger.WithContext(c)
			
			defer func() {
					stop := time.Now()

					logUnwrap := logger.Unwrap()
					logUnwrap.Info().
							Str("Service", ctx.Param("service")).
							Str("Method", ctx.Param("method")).
							// Str("remote_ip", ctx.ClientIP()).
							// Str("user_agent", string(ctx.UserAgent())).
							// Int("status", ctx.Response.StatusCode()).
							Dur("latency", stop.Sub(start)).
							Msg("request processed")
			}()
			
			ctx.Next(c)
	}
}