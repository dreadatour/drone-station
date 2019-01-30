package dshttp

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingHTTPHandler is middleware for logging HTTP requests
func LoggingHTTPHandler(next http.Handler, logger *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.WithField("duration", time.Since(start)).Infof("%s %s", r.Method, r.RequestURI)
	})
}
