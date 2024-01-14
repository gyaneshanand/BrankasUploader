package handler

import (
	"brankasv1/service"
	"encoding/json"
	"net/http"
)

// Document Upload Handler
func DocumentUploadHandler(w http.ResponseWriter, r *http.Request) {

	// Do Authentication check -> Later move it to a middleware
	// Since there is only 1 API now so writing here instead of middleware creation.
	err := service.AuthenticateRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	document, err := service.ProcessDocument(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(document)
}
