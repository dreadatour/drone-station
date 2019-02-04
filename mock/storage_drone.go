package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/pkg/dsgeo"
	"github.com/dreadatour/drone-station/storage"
)

// DroneStorage mocked
type DroneStorage struct {
	mock.Mock
}

// Check if DroneStorage mock implements storage.DroneStorage interface
var _ storage.DroneStorage = &DroneStorage{}

// List mocked
func (m *DroneStorage) List(ctx context.Context) []model.Drone {
	return m.Called().Get(0).([]model.Drone)
}

// ListWithinQuadrant mocked
func (m *DroneStorage) ListWithinQuadrant(ctx context.Context, quadrant dsgeo.Quadrant) []model.Drone {
	return m.Called(quadrant).Get(0).([]model.Drone)
}

// Add mocked
func (m *DroneStorage) Add(ctx context.Context, drone model.Drone) *model.Drone {
	return m.Called(drone).Get(0).(*model.Drone)
}

// Update mocked
func (m *DroneStorage) Update(ctx context.Context, drone model.Drone) bool {
	return m.Called(drone).Get(0).(bool)
}

// Remove mocked
func (m *DroneStorage) Remove(ctx context.Context, droneID string) bool {
	return m.Called(droneID).Get(0).(bool)
}
