package storage

import (
	"sync"

	"github.com/dreadatour/drone-station/pkg/dsgeo"

	"github.com/dreadatour/drone-station/model"
	uuid "github.com/satori/go.uuid"
)

// DroneStorage is implementation of drones dumb in-memory "database"
type DroneStorage interface {
	List() []model.Drone
	ListWithinQuadrant(quadrant dsgeo.Quadrant) []model.Drone
	Add(drone model.Drone) *model.Drone
	Remove(droneID string) bool
}

// NewDronesStorage returns new drones storage
func NewDronesStorage() DroneStorage {
	return &droneStorage{
		m: make([]model.Drone, 0),
	}
}

// droneStorage storage
type droneStorage struct {
	mx sync.RWMutex
	m  []model.Drone
}

// List of all drones
func (s *droneStorage) List() []model.Drone {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.m
}

// List of drones within quadrant
func (s *droneStorage) ListWithinQuadrant(quadrant dsgeo.Quadrant) []model.Drone {
	var res []model.Drone

	s.mx.RLock()
	defer s.mx.RUnlock()

	for _, d := range s.m {
		contains, err := quadrant.Contains(d.Latitude, d.Longitude)
		if err != nil {
			panic(err)
		}
		if contains {
			res = append(res, d)
		}
	}

	return res
}

// Add drone to storage
func (s *droneStorage) Add(drone model.Drone) *model.Drone {
	drone.ID = uuid.NewV4().String()

	s.mx.Lock()
	defer s.mx.Unlock()

	s.m = append(s.m, drone)

	return &drone
}

// Remove drone from storage
func (s *droneStorage) Remove(droneID string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()

	for i, d := range s.m {
		if d.ID == droneID {
			s.m = append(s.m[:i], s.m[i+1:]...)
			return true
		}
	}

	return false
}
