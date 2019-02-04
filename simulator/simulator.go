package simulator

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/storage"
)

// DroneSimulator is drone simulator
type DroneSimulator interface {
	Run()
}

// NewDroneSimulator returns initialised drone simulator
func NewDroneSimulator(droneStorage storage.DroneStorage, logger *logrus.Logger) DroneSimulator {
	return &droneSimulator{
		droneStorage: droneStorage,
		logger:       logger,
	}
}

type droneSimulator struct {
	droneStorage storage.DroneStorage
	logger       *logrus.Logger
}

func (s *droneSimulator) Run() {
	go s.heartbeat()
}

func (s *droneSimulator) heartbeat() {
	for {
		drones := s.droneStorage.List(context.Background())
		for _, drone := range drones {
			s.moveDrone(drone)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *droneSimulator) moveDrone(drone model.Drone) {
	phi := drone.Direction * 2 * math.Pi

	drone.Longitude += math.Cos(phi) * drone.Speed
	if drone.Longitude > 180 {
		drone.Longitude = 180 - drone.Longitude
	} else if drone.Longitude < -180 {
		drone.Longitude = 180 + drone.Longitude
	}

	drone.Latitude += math.Cos(phi) * drone.Speed
	if drone.Latitude > 90 {
		drone.Latitude = 180 - drone.Latitude
	} else if drone.Latitude < -90 {
		drone.Latitude = 180 + drone.Latitude
	}

	drone.Direction += (rand.Float64() - .5) / 10
	if drone.Direction < 0 {
		drone.Direction = 1 - drone.Direction
	} else if drone.Direction >= 1 {
		drone.Direction = drone.Direction - 1
	}

	drone.Speed += (rand.Float64() - .5) / 500000
	if drone.Speed < 0 {
		drone.Speed = 0
	}

	s.droneStorage.Update(context.Background(), drone)
}
