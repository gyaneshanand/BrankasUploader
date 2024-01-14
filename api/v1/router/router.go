package router

import (
	"brankasv1/api/v1/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// Create a MUX router
	router := mux.NewRouter()

	// Handlers for Health - Check routes
	router.HandleFunc("/api/v1/health_check", handler.HealthCheck).Methods("GET")

	router.HandleFunc("/", handler.ServeForm).Methods("GET")

	// Handler for File Upload
	router.HandleFunc("/api/v1/upload", handler.DocumentUploadHandler).Methods("POST")

	return router
}
