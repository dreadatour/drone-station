package storage

import (
	"sync"

	"github.com/dreadatour/drone-station/object"
	uuid "github.com/satori/go.uuid"
)

// Drones storage
type Drones struct {
	mx sync.RWMutex
	m  map[string]*object.Drone
}

// NewDronesStorage returns new drones storage initialised by list of drones
func NewDronesStorage(list []*object.Drone) *Drones {
	var m = make(map[string]*object.Drone)
	for _, d := range list {
		m[d.ID] = d
	}

	return &Drones{m: m}
}

// List all drones
func (s *Drones) List() []*object.Drone {
	var res []*object.Drone

	s.mx.RLock()
	defer s.mx.RUnlock()

	for _, d := range s.m {
		res = append(res, d)
	}

	return res
}

// Get drone by ID
func (s *Drones) Get(key string) (*object.Drone, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	val, ok := s.m[key]

	return val, ok
}

// Add drone to storage
func (s *Drones) Add(value object.Drone) *object.Drone {
	value.ID = uuid.NewV4().String()

	s.mx.Lock()
	defer s.mx.Unlock()

	s.m[value.ID] = &value

	return &value
}

// Remove drone from storage
func (s *Drones) Remove(ID string) bool {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.m[ID]; !ok {
		return false
	}

	delete(s.m, ID)

	return true
}
