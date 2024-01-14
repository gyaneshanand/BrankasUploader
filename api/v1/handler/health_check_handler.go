package handler

import (
	"encoding/json"
	"net/http"
)

// Health Check API Handler
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthCheckResponse := HealthCheckResponse{
		Status:  "OK",
		Message: "Brankas Server is up and running.",
	}

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(healthCheckResponse)
}

// Response from Health Check
// This type can be moved to a separate file also.
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
