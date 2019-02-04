package object_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/object"
)

func TestDroneAddValidation(t *testing.T) {
	tests := map[string]struct {
		description    string
		drone          object.DroneAdd
		expectedErrors map[string]string
	}{
		"Given empty drone object": {
			description: "All validation errors should be returned",
			drone:       object.DroneAdd{},
			expectedErrors: map[string]string{
				`quadrant`: "drone quadrant is empty",
				`x`:        "drone coordinate is empty",
				`y`:        "drone coordinate is empty",
			},
		},
		"Given drone object with incorrect coordinates X and Y": {
			description: "Coordinate validation error should be returned",
			drone:       object.DroneAdd{Quadrant: "u15pmus9", X: "foobar", Y: "121.13"},
			expectedErrors: map[string]string{
				`x`: "drone coordinate should be float from 0 to 100",
				`y`: "drone coordinate should be float from 0 to 100",
			},
		},
		"Given drone object with incorrect coordinate X": {
			description: "Coordinate validation error should be returned",
			drone:       object.DroneAdd{Quadrant: "u15pmus9", X: "foobar", Y: "12.13"},
			expectedErrors: map[string]string{
				`x`: "drone coordinate should be float from 0 to 100",
			},
		},
		"Given drone object with wrong coordinate X": {
			description: "Coordinate validation error should be returned",
			drone:       object.DroneAdd{Quadrant: "u15pmus9", X: "-1", Y: "12.13"},
			expectedErrors: map[string]string{
				`x`: "drone coordinate should be float from 0 to 100",
			},
		},
		"Given drone object with incorrect coordinate Y": {
			description: "Coordinate validation error should be returned",
			drone:       object.DroneAdd{Quadrant: "u15pmus9", X: "12.13", Y: "foobar"},
			expectedErrors: map[string]string{
				`y`: "drone coordinate should be float from 0 to 100",
			},
		},
		"Given drone object with wrong coordinate Y": {
			description: "Coordinate validation error should be returned",
			drone:       object.DroneAdd{Quadrant: "u15pmus9", X: "12.13", Y: "121.13"},
			expectedErrors: map[string]string{
				`y`: "drone coordinate should be float from 0 to 100",
			},
		},
		"Given correct drone object": {
			description:    "No validation errors should be returned",
			drone:          object.DroneAdd{Quadrant: "u15pmus9", X: "12.13", Y: "12.13"},
			expectedErrors: nil,
		},
	}
	for name, test := range tests {
		Convey(name, t, func() {
			Convey(test.description, func() {
				if test.expectedErrors == nil {
					So(test.drone.Validate(), ShouldBeNil)
				} else {
					So(test.drone.Validate(), ShouldResemble, test.expectedErrors)
				}
			})
		})
	}
}
