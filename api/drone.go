package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/dreadatour/drone-station/object"
	"github.com/dreadatour/drone-station/pkg/dshttp"
	"github.com/dreadatour/drone-station/service"
)

// DroneHandlers is drone handlers interface
type DroneHandlers interface {
	List() http.HandlerFunc
	Add() http.HandlerFunc
	Remove() http.HandlerFunc
}

// NewDroneHandlers returns initialised drone handlers interface
func NewDroneHandlers(droneService service.DroneService, logger *logrus.Logger) DroneHandlers {
	return &droneHandler{
		droneService: droneService,
		logger:       logger,
	}
}

type droneHandler struct {
	droneService service.DroneService
	logger       *logrus.Logger
}

func (h *droneHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			h.logger.WithError(err).Error("Failed to parse GET-params from request")
			dshttp.FailResponse(w, "bad-request", "Request is incorrect", nil)
			return
		}

		quadrant := r.Form.Get(`quadrant`)
		if quadrant == "" {
			dshttp.FailResponse(w, "bad-request", "Please, specify quadrant", nil)
			return
		}

		drones, err := h.droneService.List(r.Context(), quadrant)
		if err != nil {
			h.logger.WithError(err).Error("Failed to get drones list from service")
			dshttp.InternalServerErrorResponse(w)
			return
		}

		dshttp.JSONResponse(w, drones, nil)
	}
}

func (h *droneHandler) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var droneData object.DroneAdd
		if err := json.NewDecoder(r.Body).Decode(&droneData); err != nil {
			h.logger.WithError(err).Warn("Failed to decode drone JSON")
			dshttp.FailResponse(w, "bad-request", "Bad JSON", nil)
			return
		}

		if errs := droneData.Validate(); errs != nil {
			dshttp.FailResponse(w, "validation", "Drone data is invalid", errs)
			return
		}

		drone, err := h.droneService.Add(r.Context(), droneData)
		if err != nil {
			h.logger.WithError(err).Error("Failed to add new drone")
			dshttp.InternalServerErrorResponse(w)
			return
		}

		dshttp.JSONResponse(w, drone, nil)
	}
}

func (h *droneHandler) Remove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		droneID := mux.Vars(r)["droneID"]
		if droneID == "" {
			h.logger.Error("Mux var 'droneID' is not set")
			dshttp.InternalServerErrorResponse(w)
			return
		}

		err := h.droneService.Remove(r.Context(), droneID)
		if err != nil {
			if err == service.ErrDroneNotFound {
				dshttp.NotFoundErrorResponse(w)
				return
			}
			h.logger.WithError(err).Error("Failed to remove drone by ID")
			dshttp.InternalServerErrorResponse(w)
			return
		}

		dshttp.JSONEmptyResponse(w)
	}
}
