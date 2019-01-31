package object

import (
	"fmt"
	"strconv"
)

var (
	// ErrDroneQuadrantEmpty is "drone quadrant is empty" error
	ErrDroneQuadrantEmpty = fmt.Errorf("drone quadrant is empty")
	// ErrDroneCoordIsEmpty is "drone coordinate is empty" error
	ErrDroneCoordIsEmpty = fmt.Errorf("drone coordinate is empty")
	// ErrBadDroneCoord is "drone coordinate should be float from 0 to 100" error
	ErrBadDroneCoord = fmt.Errorf("drone coordinate should be float from 0 to 100")
)

// DroneAdd is request for add new drone
type DroneAdd struct {
	Quadrant string `json:"quadrant"`
	X        string `json:"x"`
	Y        string `json:"y"`
}

// Validate DroneAdd object
func (d *DroneAdd) Validate() map[string]string {
	errors := map[string]string{}

	if d.Quadrant == "" {
		errors[`quadrant`] = ErrDroneQuadrantEmpty.Error()
	}

	if d.X == "" {
		errors[`x`] = ErrDroneCoordIsEmpty.Error()
	} else {
		x, err := strconv.ParseFloat(d.X, 32)
		if err != nil || x < 0 || x > 100 {
			errors[`x`] = ErrBadDroneCoord.Error()
		}
	}

	if d.Y == "" {
		errors[`y`] = ErrDroneCoordIsEmpty.Error()
	} else {
		y, err := strconv.ParseFloat(d.Y, 32)
		if err != nil || y < 0 || y > 100 {
			errors[`y`] = ErrBadDroneCoord.Error()
		}
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}
