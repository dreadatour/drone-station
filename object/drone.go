package object

import (
	"fmt"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/pkg/dsgeo"
)

// Drone is response with drone info
type Drone struct {
	ID       string `json:"id"`
	Quadrant string `json:"quadrant"`
	X        string `json:"x"`
	Y        string `json:"y"`
}

// DroneFromModel returns a Drone object from Drone model
func DroneFromModel(drone *model.Drone, quadrant dsgeo.Quadrant) (*Drone, error) {
	x, y, err := quadrant.AbsToRel(drone.Latitude, drone.Longitude)
	if err != nil {
		return nil, err
	}

	return &Drone{
		ID:       drone.ID,
		Quadrant: quadrant.String(),
		X:        fmt.Sprintf("%.2f", x),
		Y:        fmt.Sprintf("%.2f", y),
	}, nil
}

// DronesListFromModel returns a list of Drone objects from Drone models
func DronesListFromModel(drones []model.Drone, quadrant dsgeo.Quadrant) ([]Drone, error) {
	var res = make([]Drone, len(drones))
	for i, d := range drones {
		drone, err := DroneFromModel(&d, quadrant)
		if err != nil {
			return nil, err
		}
		res[i] = *drone
	}

	return res, nil
}
