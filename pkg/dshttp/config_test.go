package dshttp_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/pkg/dshttp"
)

func TestConfig(t *testing.T) {
	Convey("Given HTTP server config with host and port", t, func() {
		c := dshttp.Config{
			Host: "0.0.0.0",
			Port: 80,
		}
		Convey("Correct address should be returned", func() {
			So(c.Address(), ShouldEqual, "0.0.0.0:80")
		})
	})
}
