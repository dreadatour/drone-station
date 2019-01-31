package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/service"
)

// DroneService mocked
type DroneService struct {
	mock.Mock
}

// Check if DroneService mock implements service.DroneService interface
var _ service.DroneService = &DroneService{}

// List mocked
func (m *DroneService) List(ctx context.Context, q string) ([]object.Drone, error) {
	args := m.Called(q)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.([]object.Drone), nil
}

// Add mocked
func (m *DroneService) Add(ctx context.Context, droneData object.DroneAdd) (*object.Drone, error) {
	args := m.Called(droneData)
	res, err := args.Get(0), args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.(*object.Drone), nil
}

// Remove mocked
func (m *DroneService) Remove(ctx context.Context, droneID string) error {
	return m.Called(droneID).Error(0)
}
