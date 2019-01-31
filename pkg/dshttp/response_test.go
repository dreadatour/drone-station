package dshttp_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/pkg/dshttp"
)

func TestInternalServerErrorResponse(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending 'Internal server error' response", func() {
			dshttp.InternalServerErrorResponse(r)

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusInternalServerError' (500)", func() {
				So(r.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Response body should be JSON with 'Internal server error' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"error","code":"internal-error","message":"Internal server error"}`)
			})
		})
	})
}

func TestNotFoundErrorResponse(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending 'Not found' response", func() {
			dshttp.NotFoundErrorResponse(r)

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusNotFound' (404)", func() {
				So(r.Code, ShouldEqual, http.StatusNotFound)
			})
			Convey("Response body should be JSON with 'Not found' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"fail","code":"not-found","message":"Not found"}`)
			})
		})
	})
}

func TestEmptyResponse(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending 'Empty' response", func() {
			dshttp.JSONEmptyResponse(r)

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("Response status code should be 'StatusOK' (200)", func() {
				So(r.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Response body should be JSON with only 'status': 'success' fields", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"success"}`)
			})
		})
	})
}

func TestFailResponse(t *testing.T) {
	tests := map[string]struct {
		description string
		code        string
		message     string
		payload     interface{}
		expected    string
	}{
		"Sending 'Fail' response without any data": {
			description: "Response body should be JSON with only 'status': 'fail' fields",
			code:        "",
			message:     "",
			payload:     nil,
			expected:    `{"status":"fail"}`,
		},
		"Sending 'Fail' response with code only": {
			description: "Response body should be JSON with code set",
			code:        "foo.bar",
			message:     "",
			payload:     nil,
			expected:    `{"status":"fail","code":"foo.bar"}`,
		},
		"Sending 'Fail' response with message only": {
			description: "Response body should be JSON with message set",
			code:        "",
			message:     "Foo bar baz",
			payload:     nil,
			expected:    `{"status":"fail","message":"Foo bar baz"}`,
		},
		"Sending 'Fail' response with code and message defined": {
			description: "Response body should be JSON with code and message set",
			code:        "foo.bar",
			message:     "Foo bar baz",
			payload:     nil,
			expected:    `{"status":"fail","code":"foo.bar","message":"Foo bar baz"}`,
		},
		"Sending 'Fail' response with payload string": {
			description: "Response body should be JSON with payload string",
			code:        "",
			message:     "",
			payload:     "foobar",
			expected:    `{"status":"fail","payload":"foobar"}`,
		},
		"Sending 'Fail' response with payload object": {
			description: "Response body should be JSON with payload object",
			code:        "",
			message:     "",
			payload:     map[string]string{"foo": "bar"},
			expected:    `{"status":"fail","payload":{"foo":"bar"}}`,
		},
		"Sending 'Fail' response with payload array": {
			description: "Response body should be JSON with payload array",
			code:        "",
			message:     "",
			payload:     []string{"foo", "bar"},
			expected:    `{"status":"fail","payload":["foo","bar"]}`,
		},
		"Sending 'Fail' response with code, message and payload object": {
			description: "Response body should be JSON with code, message and payload object",
			code:        "foo.bar",
			message:     "Foo bar baz",
			payload:     map[string]interface{}{"foo": "bar"},
			expected:    `{"status":"fail","payload":{"foo":"bar"},"code":"foo.bar","message":"Foo bar baz"}`,
		},
	}
	for name, test := range tests {
		Convey("Given HTTP response recorder", t, func() {
			r := httptest.NewRecorder()

			Convey(name, func() {
				dshttp.FailResponse(r, test.code, test.message, test.payload)

				Convey("'Content-Type' header should be equal to 'application/json'", func() {
					So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
				})
				Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
					So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
				})
				Convey("Response status code should be 'StatusBadRequest' (400)", func() {
					So(r.Code, ShouldEqual, http.StatusBadRequest)
				})
				Convey(test.description, func() {
					So(r.Body.String(), ShouldEqual, test.expected)
				})
			})
		})
	}
}

func TestFailResponseError(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending 'Fail' response with payload can't me marshalled to JSON", func() {
			dshttp.FailResponse(r, "", "", make(chan int))

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusInternalServerError' (500)", func() {
				So(r.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Response body should be JSON with 'Internal server error' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"error","code":"internal-error","message":"Internal server error"}`)
			})
		})
	})
}

func TestErrorResponse(t *testing.T) {
	tests := map[string]struct {
		description string
		code        string
		message     string
		payload     interface{}
		httpCode    int
		expected    string
	}{
		"Sending 'Error' response without any data": {
			description: "Response body should be JSON with only 'status': 'error' fields",
			code:        "",
			message:     "",
			payload:     nil,
			httpCode:    http.StatusTeapot,
			expected:    `{"status":"error"}`,
		},
		"Sending 'Error' response with code only": {
			description: "Response body should be JSON with code set",
			code:        "foo.bar",
			message:     "",
			payload:     nil,
			httpCode:    http.StatusAlreadyReported,
			expected:    `{"status":"error","code":"foo.bar"}`,
		},
		"Sending 'Error' response with message only": {
			description: "Response body should be JSON with message set",
			code:        "",
			message:     "Foo bar baz",
			payload:     nil,
			httpCode:    http.StatusExpectationFailed,
			expected:    `{"status":"error","message":"Foo bar baz"}`,
		},
		"Sending 'Error' response with code and message defined": {
			description: "Response body should be JSON with code and message set",
			code:        "foo.bar",
			message:     "Foo bar baz",
			payload:     nil,
			httpCode:    http.StatusHTTPVersionNotSupported,
			expected:    `{"status":"error","code":"foo.bar","message":"Foo bar baz"}`,
		},
		"Sending 'Error' response with payload string": {
			description: "Response body should be JSON with payload string",
			message:     "",
			payload:     "foobar",
			httpCode:    http.StatusConflict,
			expected:    `{"status":"error","payload":"foobar"}`,
		},
		"Sending 'Error' response with payload object": {
			description: "Response body should be JSON with payload object",
			code:        "",
			message:     "",
			payload:     map[string]string{"foo": "bar"},
			httpCode:    http.StatusVariantAlsoNegotiates,
			expected:    `{"status":"error","payload":{"foo":"bar"}}`,
		},
		"Sending 'Error' response with payload array": {
			description: "Response body should be JSON with payload array",
			code:        "",
			message:     "",
			payload:     []string{"foo", "bar"},
			httpCode:    http.StatusNetworkAuthenticationRequired,
			expected:    `{"status":"error","payload":["foo","bar"]}`,
		},
		"Sending 'Error' response with code, message and payload object": {
			description: "Response body should be JSON with code, message and payload object",
			code:        "foo.bar",
			message:     "Foo bar baz",
			payload:     map[string]interface{}{"foo": "bar"},
			httpCode:    http.StatusTooManyRequests,
			expected:    `{"status":"error","payload":{"foo":"bar"},"code":"foo.bar","message":"Foo bar baz"}`,
		},
	}
	for name, test := range tests {
		Convey("Given HTTP response recorder", t, func() {
			r := httptest.NewRecorder()

			Convey(name, func() {
				dshttp.ErrorResponse(r, test.code, test.message, test.payload, test.httpCode)

				Convey("'Content-Type' header should be equal to 'application/json'", func() {
					So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
				})
				Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
					So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
				})
				Convey("Response status code should be correct", func() {
					So(r.Code, ShouldEqual, test.httpCode)
				})
				Convey(test.description, func() {
					So(r.Body.String(), ShouldEqual, test.expected)
				})
			})
		})
	}
}

func TestErrorResponseError(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending 'Error' response with payload can't me marshalled to JSON", func() {
			dshttp.ErrorResponse(r, "", "", make(chan int), http.StatusTeapot)

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusInternalServerError' (500)", func() {
				So(r.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Response body should be JSON with 'Internal server error' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"error","code":"internal-error","message":"Internal server error"}`)
			})
		})
	})
}

func TestJSONResponse(t *testing.T) {
	tests := map[string]struct {
		description string
		payload     interface{}
		meta        interface{}
		expected    string
	}{
		"Sending JSON response without any data": {
			description: "Response body should be JSON with only 'status': 'success' fields",
			payload:     nil,
			meta:        nil,
			expected:    `{"status":"success"}`,
		},
		"Sending JSON response with payload array": {
			description: "Response body should be JSON with payload array and empty meta",
			payload:     []string{"foo", "bar"},
			meta:        nil,
			expected:    `{"status":"success","payload":["foo","bar"]}`,
		},
		"Sending JSON response with payload object": {
			description: "Response body should be JSON with payload object and empty meta",
			payload:     map[string]interface{}{"foo": "bar"},
			meta:        nil,
			expected:    `{"status":"success","payload":{"foo":"bar"}}`,
		},
		"Sending JSON response with meta object": {
			description: "Response body should be JSON with empty payload and meta object",
			payload:     nil,
			meta:        map[string]interface{}{"foo": "bar"},
			expected:    `{"status":"success","meta":{"foo":"bar"}}`,
		},
		"Sending JSON response with payload object and meta string": {
			description: "Response body should be JSON with payload object and meta string",
			payload:     map[string]interface{}{"foo": "bar"},
			meta:        "foobar",
			expected:    `{"status":"success","payload":{"foo":"bar"},"meta":"foobar"}`,
		},
		"Sending JSON response with payload and meta objects": {
			description: "Response body should be JSON with payload and meta objects",
			payload:     map[string]interface{}{"foo": "bar"},
			meta:        map[string]int{"foo": 1},
			expected:    `{"status":"success","payload":{"foo":"bar"},"meta":{"foo":1}}`,
		},
	}
	for name, test := range tests {
		Convey("Given HTTP response recorder", t, func() {
			r := httptest.NewRecorder()

			Convey(name, func() {
				dshttp.JSONResponse(r, test.payload, test.meta)

				Convey("'Content-Type' header should be equal to 'application/json'", func() {
					So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
				})
				Convey("Response status code should be 'StatusOK' (200)", func() {
					So(r.Code, ShouldEqual, http.StatusOK)
				})
				Convey(test.description, func() {
					So(r.Body.String(), ShouldEqual, test.expected)
				})
			})
		})
	}
}

func TestJSONResponseErrorPayload(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending JSON response with payload can't me marshalled to JSON", func() {
			dshttp.JSONResponse(r, make(chan int), nil)

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusInternalServerError' (500)", func() {
				So(r.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Response body should be JSON with 'Internal server error' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"error","code":"internal-error","message":"Internal server error"}`)
			})
		})
	})
}

func TestJSONResponseErrorMeta(t *testing.T) {
	Convey("Given HTTP response recorder", t, func() {
		r := httptest.NewRecorder()

		Convey("Sending JSON response with meta can't me marshalled to JSON", func() {
			dshttp.JSONResponse(r, nil, make(chan int))

			Convey("'Content-Type' header should be equal to 'application/json'", func() {
				So(r.Header().Get("Content-Type"), ShouldEqual, "application/json")
			})
			Convey("'X-Content-Type-Options' header should be equal to 'nosniff'", func() {
				So(r.Header().Get("X-Content-Type-Options"), ShouldEqual, "nosniff")
			})
			Convey("Response status code should be 'StatusInternalServerError' (500)", func() {
				So(r.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Response body should be JSON with 'Internal server error' error", func() {
				So(r.Body.String(), ShouldEqual, `{"status":"error","code":"internal-error","message":"Internal server error"}`)
			})
		})
	})
}
