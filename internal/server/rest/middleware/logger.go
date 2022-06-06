package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/romangurevitch/golang-concurrency/internal/pkg/logger"
	"github.com/sirupsen/logrus"
	"time"
)

// LogMiddleware logs a gin HTTP requests in JSON format, with some additional custom key/values
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.WithContext(c.Request.Context())
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		entry := log.WithFields(logrus.Fields{
			"duration":   time.Since(start).Milliseconds(),
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
