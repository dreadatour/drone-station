package storage

import (
	"context"
	"sync"

	"github.com/dreadatour/drone-station/pkg/dsgeo"
	"github.com/sirupsen/logrus"

	"github.com/dreadatour/drone-station/model"
	uuid "github.com/satori/go.uuid"
)

// DroneStorage is implementation of drones dumb in-memory "database"
type DroneStorage interface {
	List(ctx context.Context) []model.Drone
	ListWithinQuadrant(ctx context.Context, quadrant dsgeo.Quadrant) []model.Drone
	Add(ctx context.Context, drone model.Drone) *model.Drone
	Remove(ctx context.Context, droneID string) bool
}

// NewDronesStorage returns new drones storage
func NewDronesStorage(logger *logrus.Logger) DroneStorage {
	return &droneStorage{
		m:      make([]model.Drone, 0),
		logger: logger,
	}
}

// droneStorage storage
type droneStorage struct {
	mx     sync.RWMutex
	m      []model.Drone
	logger *logrus.Logger
}

// List of all drones
func (s *droneStorage) List(ctx context.Context) []model.Drone {
	s.mx.RLock()
	defer s.mx.RUnlock()

	return s.m
}

// List of drones within quadrant
func (s *droneStorage) ListWithinQuadrant(ctx context.Context, quadrant dsgeo.Quadrant) []model.Drone {
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
func (s *droneStorage) Add(ctx context.Context, drone model.Drone) *model.Drone {
	drone.ID = uuid.NewV4().String()

	s.mx.Lock()
	defer s.mx.Unlock()

	s.m = append(s.m, drone)

	return &drone
}

// Remove drone from storage
func (s *droneStorage) Remove(ctx context.Context, droneID string) bool {
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
