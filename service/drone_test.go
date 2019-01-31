package service_test

import (
	"context"
	"testing"

	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/pkg/dsgeo"

	"github.com/sirupsen/logrus/hooks/test"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/mock"
	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/service"
)

func TestNewDroneService(t *testing.T) {
	Convey("Creating new drone service", t, func() {
		storage := new(mock.DroneStorage)
		logger, _ := test.NewNullLogger()

		droneService := service.NewDroneService(storage, logger)

		Convey("Service should be created", func() {
			So(droneService, ShouldNotBeNil)
		})
	})
}

func TestDroneServiceList(t *testing.T) {
	Convey("Having drone service", t, func() {
		storage := new(mock.DroneStorage)
		logger, _ := test.NewNullLogger()
		droneService := service.NewDroneService(storage, logger)

		Convey("Having no drones in storage", func() {
			storage.On("ListWithinQuadrant", dsgeo.NewQuadrantFromString("u15pmus9")).Return([]model.Drone{})

			Convey("List response should be empty", func() {
				drones, err := droneService.List(context.Background(), "u15pmus9")
				So(err, ShouldBeNil)
				So(len(drones), ShouldEqual, 0)

				So(storage.AssertExpectations(t), ShouldBeTrue)
			})
		})

		Convey("Having drones in storage", func() {
			storage.On("ListWithinQuadrant", dsgeo.NewQuadrantFromString("u15pmus9")).Return([]model.Drone{
				model.Drone{
					ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
					Latitude:  51.92442,
					Longitude: 4.477736,
				},
			})

			Convey("List response should contains drone from storage", func() {
				drones, err := droneService.List(context.Background(), "u15pmus9")
				So(err, ShouldBeNil)
				So(len(drones), ShouldEqual, 1)
				So(drones[0], ShouldResemble, object.Drone{
					ID:       "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
					Quadrant: "u15pmus9",
					X:        "35.14",
					Y:        "67.01",
				})

				So(storage.AssertExpectations(t), ShouldBeTrue)
			})
		})
	})
}

func TestDroneServiceAdd(t *testing.T) {
	Convey("Having drone service", t, func() {
		storage := new(mock.DroneStorage)
		logger, _ := test.NewNullLogger()
		droneService := service.NewDroneService(storage, logger)

		Convey("Adding new drone", func() {
			storage.On("Add", model.Drone{Latitude: 51.92436332702637, Longitude: 4.477656555175781}).Return(&model.Drone{ID: "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e", Latitude: 51.92436332702637, Longitude: 4.477656555175781})

			drone, err := droneService.Add(context.Background(), object.DroneAdd{Quadrant: "u15pmus9", X: "12", Y: "34"})

			Convey("Drone added should be returned", func() {
				So(err, ShouldBeNil)
				So(drone.ID, ShouldEqual, "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e")
				So(drone.Quadrant, ShouldEqual, "u15pmus9")
				So(drone.X, ShouldEqual, "12.00")
				So(drone.Y, ShouldEqual, "34.00")
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})

		Convey("Adding new drone with incorrect X coordinate", func() {
			drone, err := droneService.Add(context.Background(), object.DroneAdd{Quadrant: "u15pmus9", X: "foo", Y: "34"})

			Convey("Error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(drone, ShouldBeNil)
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})

		Convey("Adding new drone with incorrect Y coordinate", func() {
			drone, err := droneService.Add(context.Background(), object.DroneAdd{Quadrant: "u15pmus9", X: "12", Y: "bar"})

			Convey("Error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(drone, ShouldBeNil)
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})

		Convey("Adding new drone with wrong X and Y coordinates", func() {
			drone, err := droneService.Add(context.Background(), object.DroneAdd{Quadrant: "u15pmus9", X: "121", Y: "-343"})

			Convey("Error should be returned", func() {
				So(err, ShouldNotBeNil)
				So(drone, ShouldBeNil)
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})
	})
}

func TestDroneServiceRemove(t *testing.T) {
	Convey("Having drone service", t, func() {
		storage := new(mock.DroneStorage)
		logger, _ := test.NewNullLogger()
		droneService := service.NewDroneService(storage, logger)

		Convey("Removing unexisting drone", func() {
			storage.On("Remove", "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e").Return(false)

			err := droneService.Remove(context.Background(), "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e")

			Convey("Error should be returned", func() {
				So(err, ShouldEqual, service.ErrDroneNotFound)
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})

		Convey("Removing existing drone", func() {
			storage.On("Remove", "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e").Return(true)

			err := droneService.Remove(context.Background(), "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e")

			Convey("No error should be returned", func() {
				So(err, ShouldBeNil)
			})

			So(storage.AssertExpectations(t), ShouldBeTrue)
		})
	})
}
