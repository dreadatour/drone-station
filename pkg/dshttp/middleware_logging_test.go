package dshttp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/pkg/dshttp"
)

func TestLoggingHTTPHandler(t *testing.T) {
	Convey("Given HTTP handler with logging middleware", t, func() {
		logger, hook := test.NewNullLogger()
		handler := dshttp.LoggingHTTPHandler(getTestHandler(), logger)

		Convey("Serving HTTP request", func() {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			So(err, ShouldBeNil)
			req.RequestURI = "/path"

			r := httptest.NewRecorder()
			handler.ServeHTTP(r, req)

			Convey("Response should be OK", func() {
				So(r.Code, ShouldEqual, http.StatusOK)
				So(r.Body.String(), ShouldEqual, "OK")
			})

			Convey("Logger should have one entry", func() {
				So(hook.Entries, ShouldHaveLength, 1)

				logEntry := hook.Entries[0]

				Convey("With 'info' level", func() {
					So(logEntry.Level, ShouldEqual, logrus.InfoLevel)
				})
				Convey("With 'GET /path' message", func() {
					So(logEntry.Message, ShouldEqual, logEntry.Message)
				})
				Convey("With 'duration' field set", func() {
					durationValue, ok := logEntry.Data["duration"]
					So(ok, ShouldBeTrue)

					Convey("With duration greater than zero", func() {
						duration, ok := durationValue.(time.Duration)
						So(ok, ShouldBeTrue)
						So(duration, ShouldBeGreaterThan, 0)
					})
				})
			})
		})
	})
}

// getTestHandler returns a http.HandlerFunc for testing http middleware
func getTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("OK"))
	}
	return http.HandlerFunc(fn)
}
