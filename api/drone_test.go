package api_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus/hooks/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/api"
	"github.com/dreadatour/drone-station/mock"
	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/pkg/dshttp"
	"github.com/dreadatour/drone-station/service"
)

func TestNewDroneHandlers(t *testing.T) {
	Convey("Creating new drone handlers", t, func() {
		droneService := new(mock.DroneService)
		logger, _ := test.NewNullLogger()

		droneHandlers := api.NewDroneHandlers(droneService, logger)

		Convey("Handlers should be created", func() {
			So(droneHandlers, ShouldNotBeNil)
		})
	})
}

func TestDroneHandlersList(t *testing.T) {
	Convey("Having drone handlers", t, func() {
		droneService := new(mock.DroneService)
		logger, _ := test.NewNullLogger()
		droneHandlers := api.NewDroneHandlers(droneService, logger)

		Convey("Given a GET request for /api/v1/drones with quadrant set", func() {
			req := httptest.NewRequest("GET", "/api/v1/drones?quadrant=u15pmus9", nil)
			resp := httptest.NewRecorder()

			Convey("Handling request by List handler with no drones found", func() {
				droneService.On("List", "u15pmus9").Return([]object.Drone{}, nil)

				droneHandlers.List()(resp, req)

				Convey("Response code should be 200 OK", func() {
					So(resp.Code, ShouldEqual, 200)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "success")
					So(body.Payload, ShouldResemble, []interface{}{})
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})

			Convey("Handling request by List handler with drones found", func() {
				droneService.On("List", "u15pmus9").Return([]object.Drone{object.Drone{ID: "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", Quadrant: "u15pmus9", X: "12.12", Y: "34.34"}}, nil)

				droneHandlers.List()(resp, req)

				Convey("Response code should be 200 OK", func() {
					So(resp.Code, ShouldEqual, 200)
				})

				Convey("Response body should have payload with drone found", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "success")
					So(body.Payload, ShouldResemble, []interface{}{map[string]interface{}{`id`: "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", `quadrant`: "u15pmus9", `x`: "12.12", `y`: "34.34"}})
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})

			Convey("Handling request by List handler with error returned by drone service", func() {
				droneService.On("List", "u15pmus9").Return(nil, fmt.Errorf("DEADBEEF"))

				droneHandlers.List()(resp, req)

				Convey("Response code should be 500 Internal Server Error", func() {
					So(resp.Code, ShouldEqual, 500)
				})

				Convey("Response body should have payload with drone found", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "error")
					So(body.Code, ShouldEqual, "internal-error")
					So(body.Message, ShouldEqual, "Internal server error")
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})
		})

		Convey("Given a GET request for /api/v1/drones with quadrant not set", func() {
			req := httptest.NewRequest("GET", "/api/v1/drones", nil)
			resp := httptest.NewRecorder()

			Convey("Handling request by List handler", func() {
				droneHandlers.List()(resp, req)

				Convey("Response code should be 400 Bad Request", func() {
					So(resp.Code, ShouldEqual, 400)
				})

				Convey("Response body should be an error", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "fail")
					So(body.Code, ShouldEqual, "bad-request")
					So(body.Message, ShouldEqual, "Please, specify quadrant")
				})
			})
		})

		Convey("Given a GET request for /api/v1/drones with wrong form", func() {
			req := httptest.NewRequest("GET", "/api/v1/drones?%a", nil)
			resp := httptest.NewRecorder()

			Convey("Handling request by List handler", func() {
				droneHandlers.List()(resp, req)

				Convey("Response code should be 400 Bad Request", func() {
					So(resp.Code, ShouldEqual, 400)
				})

				Convey("Response body should be an error", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "fail")
					So(body.Code, ShouldEqual, "bad-request")
					So(body.Message, ShouldEqual, "Request is incorrect")
				})
			})
		})
	})
}

func TestDroneHandlersAdd(t *testing.T) {
	Convey("Having drone handlers", t, func() {
		droneService := new(mock.DroneService)
		logger, _ := test.NewNullLogger()
		droneHandlers := api.NewDroneHandlers(droneService, logger)

		Convey("Given a POST request for /api/v1/drones with valid data", func() {
			req := httptest.NewRequest("POST", "/api/v1/drones", strings.NewReader(`{"quadrant": "u15pmus9", "x":"12.12", "y": "34.34"}`))
			resp := httptest.NewRecorder()

			Convey("Handling request by Add handler", func() {
				droneService.On("Add", object.DroneAdd{Quadrant: "u15pmus9", X: "12.12", Y: "34.34"}).Return(&object.Drone{ID: "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", Quadrant: "u15pmus9", X: "12.12", Y: "34.34"}, nil)

				droneHandlers.Add()(resp, req)

				Convey("Response code should be 200 OK", func() {
					So(resp.Code, ShouldEqual, 200)
				})

				Convey("Response body should have payload with drone found", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "success")
					So(body.Payload, ShouldResemble, map[string]interface{}{`id`: "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", `quadrant`: "u15pmus9", `x`: "12.12", `y`: "34.34"})
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})

			Convey("Handling request by Add handler with error returned by drone service", func() {
				droneService.On("Add", object.DroneAdd{Quadrant: "u15pmus9", X: "12.12", Y: "34.34"}).Return(nil, fmt.Errorf("DEADBEEF"))

				droneHandlers.Add()(resp, req)

				Convey("Response code should be 500 Internal Server Error", func() {
					So(resp.Code, ShouldEqual, 500)
				})

				Convey("Response body should have payload with drone found", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "error")
					So(body.Code, ShouldEqual, "internal-error")
					So(body.Message, ShouldEqual, "Internal server error")
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})
		})

		Convey("Given a POST request for /api/v1/drones with no data", func() {
			req := httptest.NewRequest("POST", "/api/v1/drones", nil)
			resp := httptest.NewRecorder()

			Convey("Handling request by Add handler", func() {
				droneHandlers.Add()(resp, req)

				Convey("Response code should be 400 Bad Request", func() {
					So(resp.Code, ShouldEqual, 400)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "fail")
					So(body.Code, ShouldEqual, "bad-request")
					So(body.Message, ShouldEqual, "Bad JSON")
				})
			})
		})

		Convey("Given a POST request for /api/v1/drones with invalid data", func() {
			req := httptest.NewRequest("POST", "/api/v1/drones", strings.NewReader(`{"x":"12.12", "y": "34.34"}`))
			resp := httptest.NewRecorder()

			Convey("Handling request by Add handler", func() {
				droneHandlers.Add()(resp, req)

				Convey("Response code should be 400 Bad Request", func() {
					So(resp.Code, ShouldEqual, 400)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "fail")
					So(body.Code, ShouldEqual, "validation")
					So(body.Message, ShouldEqual, "Drone data is invalid")
					So(body.Payload, ShouldResemble, map[string]interface{}{`quadrant`: "drone quadrant is empty"})
				})
			})
		})
	})
}

func TestDroneHandlersRemove(t *testing.T) {
	Convey("Having drone handlers", t, func() {
		droneService := new(mock.DroneService)
		logger, _ := test.NewNullLogger()
		droneHandlers := api.NewDroneHandlers(droneService, logger)

		Convey("Given a DELETE request for /api/v1/drones/{droneID}", func() {
			req := httptest.NewRequest("DELETE", "/api/v1/drones/45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", nil)
			req = mux.SetURLVars(req, map[string]string{"droneID": "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e"})
			resp := httptest.NewRecorder()

			Convey("Handling request by Remove handler", func() {
				droneService.On("Remove", "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e").Return(nil)

				droneHandlers.Remove()(resp, req)

				Convey("Response code should be 200 OK", func() {
					So(resp.Code, ShouldEqual, 200)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "success")
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})

			Convey("Handling request by Remove handler when drone is not found", func() {
				droneService.On("Remove", "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e").Return(service.ErrDroneNotFound)

				droneHandlers.Remove()(resp, req)

				Convey("Response code should be 404 Not Found", func() {
					So(resp.Code, ShouldEqual, 404)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "fail")
					So(body.Code, ShouldEqual, "not-found")
					So(body.Message, ShouldEqual, "Not found")
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})

			Convey("Handling request by Remove handler with error returned by drone service", func() {
				droneService.On("Remove", "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e").Return(fmt.Errorf("DEADBEEF"))

				droneHandlers.Remove()(resp, req)

				Convey("Response code should be 500 Internal Server Error", func() {
					So(resp.Code, ShouldEqual, 500)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "error")
					So(body.Code, ShouldEqual, "internal-error")
					So(body.Message, ShouldEqual, "Internal server error")
				})

				So(droneService.AssertExpectations(t), ShouldBeTrue)
			})
		})

		Convey("Given a DELETE request for /api/v1/drones/{droneID} with droneID mux var not set", func() {
			req := httptest.NewRequest("DELETE", "/api/v1/drones/45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", nil)
			resp := httptest.NewRecorder()

			Convey("Handling request by Remove handler", func() {
				droneHandlers.Remove()(resp, req)

				Convey("Response code should be 500 Internal Server Error", func() {
					So(resp.Code, ShouldEqual, 500)
				})

				Convey("Response body should have empty payload", func() {
					body := dshttp.Response{}
					json.Unmarshal(resp.Body.Bytes(), &body)

					So(body.Status, ShouldEqual, "error")
					So(body.Code, ShouldEqual, "internal-error")
					So(body.Message, ShouldEqual, "Internal server error")
				})
			})
		})
	})
}
