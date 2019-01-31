package dslog_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/pkg/dslog"
)

func TestLevel(t *testing.T) {
	tests := map[string]struct {
		description   string
		level         string
		expected      logrus.Level
		errorExpected bool
	}{
		"Given empty log level": {
			description:   "Error should be returned",
			level:         "",
			expected:      0,
			errorExpected: true,
		},
		"Given 'debug' log level": {
			description:   "Logrus 'DebugLevel' should be used",
			level:         "debug",
			expected:      logrus.DebugLevel,
			errorExpected: false,
		},
		"Given 'info' log level": {
			description:   "Logrus 'InfoLevel' should be used",
			level:         "info",
			expected:      logrus.InfoLevel,
			errorExpected: false,
		},
		"Given 'warning' log level": {
			description:   "Logrus 'WarnLevel' should be used",
			level:         "warning",
			expected:      logrus.WarnLevel,
			errorExpected: false,
		},
		"Given 'error' log level": {
			description:   "Logrus 'ErrorLevel' should be used",
			level:         "error",
			expected:      logrus.ErrorLevel,
			errorExpected: false,
		},
		"Given 'fatal' log level": {
			description:   "Logrus 'FatalLevel' should be used",
			level:         "fatal",
			expected:      logrus.FatalLevel,
			errorExpected: false,
		},
		"Given 'panic' log level": {
			description:   "Logrus 'PanicLevel' should be used",
			level:         "panic",
			expected:      logrus.PanicLevel,
			errorExpected: false,
		},
		"Given unknown log level": {
			description:   "Error should be returned",
			level:         "foobar",
			expected:      0,
			errorExpected: true,
		},
	}
	for name, test := range tests {
		Convey(name, t, func() {
			logger, err := dslog.NewLogger(&dslog.Config{Level: test.level, Format: "text"})
			Convey(test.description, func() {
				if test.errorExpected {
					So(err, ShouldBeError)
					So(logger, ShouldBeNil)
				} else {
					So(logger.Level, ShouldEqual, test.expected)
					So(err, ShouldBeNil)
				}
			})
		})
	}
}

func TestFormat(t *testing.T) {
	tests := map[string]struct {
		description   string
		format        string
		errorExpected bool
	}{
		"Given empty log format": {
			description:   "Error should be returned",
			format:        "",
			errorExpected: true,
		},
		"Given 'text' log format": {
			description:   "Logrus 'TextFormatter' should be used",
			format:        "text",
			errorExpected: false,
		},
		"Given 'plaintext' log format": {
			description:   "Logrus 'TextFormatter' should be used",
			format:        "plaintext",
			errorExpected: false,
		},
		"Given 'json' log format": {
			description:   "Logrus 'JSONFormatter' should be used",
			format:        "json",
			errorExpected: false,
		},
		"Given unknown log format": {
			description:   "Error should be returned",
			format:        "foobar",
			errorExpected: true,
		},
	}
	for name, test := range tests {
		Convey(name, t, func() {
			_, err := dslog.NewLogger(&dslog.Config{Level: "info", Format: test.format})
			Convey(test.description, func() {
				if test.errorExpected {
					So(err, ShouldBeError)
				} else {
					So(err, ShouldBeNil)
				}
			})
		})
	}
}
