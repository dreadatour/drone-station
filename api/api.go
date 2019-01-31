package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	"github.com/dreadatour/drone-station/pkg/dshttp"
	"github.com/dreadatour/drone-station/pkg/dslog"
	"github.com/dreadatour/drone-station/service"
	"github.com/dreadatour/drone-station/storage"
)

// Config server
type Config struct {
	HTTP *dshttp.Config `envconfig:"HTTP"`
	Log  *dslog.Config  `envconfig:"LOG"`
}

// Run API server
func Run() error {
	// parse config
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		envconfig.Usage("", &cfg)
		return err
	}

	// initialise logger
	logger, err := dslog.NewLogger(cfg.Log)
	if err != nil {
		return err
	}

	// storage
	var droneStorage = storage.NewDronesStorage(logger)

	// service
	var droneService = service.NewDroneService(droneStorage, logger)

	// handlers
	var droneHandlers = NewDroneHandlers(droneService, logger)

	// initialise router
	router := mux.NewRouter().StrictSlash(true)

	// user routes
	router.Path("/api/v1/drones").Methods(http.MethodGet).HandlerFunc(droneHandlers.List())
	router.Path("/api/v1/drones").Methods(http.MethodPost).HandlerFunc(droneHandlers.Add())
	router.Path("/api/v1/drones/{droneID}").Methods(http.MethodDelete).HandlerFunc(droneHandlers.Remove())

	// start HTTP server
	addr := cfg.HTTP.Address()
	logger.Infof("Starting HTTP server at %s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: dshttp.LoggingHTTPHandler(router, logger),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				logger.WithError(err).Error("Error starting HTTP server")
			}
		}
	}()

	// wait for a signal to quit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	// graceful shutdown
	cleanupDone := make(chan struct{})
	go func() {
		<-signalChan

		logger.Info("Received an interrupt, shutdown HTTP server")
		ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctxCancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.WithError(err).Error("Error stopping HTTP server")
		}

		close(cleanupDone)
	}()

	// wait until cleanup done (successfully or not) before exit
	<-cleanupDone
	return nil
}
