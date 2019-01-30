package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/pkg/dshttp"
	"github.com/dreadatour/drone-station/storage"
)

// DroneHandlers is drone handlers interface
type DroneHandlers interface {
	List() http.HandlerFunc
	Create() http.HandlerFunc
	Delete() http.HandlerFunc
}

// NewDroneHandlers returns initialised drone handlers interface
func NewDroneHandlers(droneStorage *storage.Drones, logger *logrus.Logger) DroneHandlers {
	return &droneHandler{
		droneStorage: droneStorage,
		logger:       logger,
	}
}

type droneHandler struct {
	droneStorage *storage.Drones
	logger       *logrus.Logger
}

func (h *droneHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		drones := h.droneStorage.List()

		dshttp.JSONResponse(w, drones, nil)
	}
}

func (h *droneHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req object.DroneAdd
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			dshttp.FailResponse(w, "bad-request", "Bad JSON", nil)
			return
		}

		newDrone := object.Drone{
			Quadrant: req.Quadrant,
			X:        req.X,
			Y:        req.Y,
		}

		drone := h.droneStorage.Add(newDrone)

		dshttp.JSONResponse(w, drone, nil)
	}
}

func (h *droneHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		droneID := mux.Vars(r)["droneID"]
		if droneID == "" {
			h.logger.Error("Mux var 'droneID' is not set")
			dshttp.InternalServerErrorResponse(w)
			return
		}

		ok := h.droneStorage.Remove(droneID)
		if !ok {
			dshttp.NotFoundErrorResponse(w)
			return
		}

		dshttp.JSONEmptyResponse(w)
	}
}
