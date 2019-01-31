package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/pkg/dsgeo"
	"github.com/dreadatour/drone-station/storage"
)

// ErrDroneNotFound is "drone not found" error
var ErrDroneNotFound = fmt.Errorf("drone not found")

// DroneService is drone service
type DroneService interface {
	List(ctx context.Context, q string) ([]object.Drone, error)
	Add(ctx context.Context, droneData object.DroneAdd) (*object.Drone, error)
	Remove(ctx context.Context, droneID string) error
}

// NewDroneService returns initialised drone service
func NewDroneService(droneStorage storage.DroneStorage, logger *logrus.Logger) DroneService {
	return &droneService{
		droneStorage: droneStorage,
		logger:       logger,
	}
}

type droneService struct {
	droneStorage storage.DroneStorage
	logger       *logrus.Logger
}

func (s *droneService) List(ctx context.Context, q string) ([]object.Drone, error) {
	quadrant := dsgeo.NewQuadrantFromString(q)

	drones := s.droneStorage.ListWithinQuadrant(ctx, quadrant)

	return object.DronesListFromModel(drones, quadrant)

}

func (s *droneService) Add(ctx context.Context, droneData object.DroneAdd) (*object.Drone, error) {
	quadrant := dsgeo.NewQuadrantFromString(droneData.Quadrant)

	x, err := strconv.ParseFloat(droneData.X, 64)
	if err != nil {
		return nil, err
	}

	y, err := strconv.ParseFloat(droneData.Y, 64)
	if err != nil {
		return nil, err
	}

	latitude, longitude, err := quadrant.RelToAbs(x, y)
	if err != nil {
		return nil, err
	}

	drone := s.droneStorage.Add(ctx, model.Drone{
		Latitude:  latitude,
		Longitude: longitude,
	})

	return object.DroneFromModel(drone, quadrant)
}

func (s *droneService) Remove(ctx context.Context, droneID string) error {
	ok := s.droneStorage.Remove(ctx, droneID)
	if !ok {
		return ErrDroneNotFound
	}

	return nil
}
