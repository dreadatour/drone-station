package object_test

import (
	"testing"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/pkg/dsgeo"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/object"
)

func TestDroneFromModel(t *testing.T) {
	tests := map[string]struct {
		description   string
		drone         *model.Drone
		quadrant      dsgeo.Quadrant
		expected      *object.Drone
		expectedError error
	}{
		"Given drone and quadrant": {
			description: "Drone object should be returned",
			drone: &model.Drone{
				ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
				Latitude:  51.924333714,
				Longitude: 4.477883541,
			},
			quadrant: dsgeo.NewQuadrantFromString("u15pmus9"),
			expected: &object.Drone{
				ID:       "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
				Quadrant: "u15pmus9",
				X:        "78.11",
				Y:        "16.75",
			},
			expectedError: nil,
		},
		"Given drone with coordinates not belongs to given quadrant": {
			description: "Drone object should be returned",
			drone: &model.Drone{
				ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
				Latitude:  51.925333714,
				Longitude: 4.473883541,
			},
			quadrant:      dsgeo.NewQuadrantFromString("u15pmus9"),
			expected:      nil,
			expectedError: dsgeo.ErrCoordinateOutOfRange,
		},
		"Given drone with wrong coordinates and quadrant": {
			description: "Drone object should be returned",
			drone: &model.Drone{
				ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
				Latitude:  200,
				Longitude: 0,
			},
			quadrant:      dsgeo.NewQuadrantFromString("u15pmus9"),
			expected:      nil,
			expectedError: dsgeo.ErrBadLatitude,
		},
	}
	for name, test := range tests {
		Convey(name, t, func() {
			Convey(test.description, func() {
				drone, err := object.DroneFromModel(test.drone, test.quadrant)
				if test.expectedError == nil {
					So(err, ShouldBeNil)
					So(drone, ShouldResemble, test.expected)
				} else {
					So(err, ShouldEqual, test.expectedError)
				}
			})
		})
	}
}

func TestDronesListFromModel(t *testing.T) {
	tests := map[string]struct {
		description   string
		drones        []model.Drone
		quadrant      dsgeo.Quadrant
		expected      []object.Drone
		expectedError error
	}{
		"Given drones list and quadrant": {
			description: "Drone objects list should be returned",
			drones: []model.Drone{
				model.Drone{
					ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
					Latitude:  51.924333714,
					Longitude: 4.477883541,
				},
				model.Drone{
					ID:        "75eb2a99-d55e-4874-9ca2-cddc3784afaf",
					Latitude:  51.924433714,
					Longitude: 4.477783541,
				},
			},
			quadrant: dsgeo.NewQuadrantFromString("u15pmus9"),
			expected: []object.Drone{
				object.Drone{
					ID:       "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
					Quadrant: "u15pmus9",
					X:        "78.11",
					Y:        "16.75",
				},
				object.Drone{
					ID:       "75eb2a99-d55e-4874-9ca2-cddc3784afaf",
					Quadrant: "u15pmus9",
					X:        "48.99",
					Y:        "75.00",
				},
			},
			expectedError: nil,
		},
		"Given drones list with coordinates not belongs to given quadrant": {
			description: "Drone objects list should be returned",
			drones: []model.Drone{
				model.Drone{
					ID:        "45745c60-7b1a-11e8-9c9c-2d42b21b1a3e",
					Latitude:  0,
					Longitude: 0,
				},
				model.Drone{
					ID:        "75eb2a99-d55e-4874-9ca2-cddc3784afaf",
					Latitude:  51.924433714,
					Longitude: 4.477783541,
				},
			},
			quadrant:      dsgeo.NewQuadrantFromString("u15pmus9"),
			expected:      nil,
			expectedError: dsgeo.ErrCoordinateOutOfRange,
		},
	}
	for name, test := range tests {
		Convey(name, t, func() {
			Convey(test.description, func() {
				drone, err := object.DronesListFromModel(test.drones, test.quadrant)
				if test.expectedError == nil {
					So(err, ShouldBeNil)
					So(drone, ShouldResemble, test.expected)
				} else {
					So(err, ShouldEqual, test.expectedError)
				}
			})
		})
	}
}
