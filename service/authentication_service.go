package service

import (
	"errors"
	"net/http"
	"os"
)

func AuthenticateRequest(r *http.Request) error {

	requestAuthToken := r.FormValue("authtoken") // Get Token from request
	authToken := os.Getenv("AUTH_TOKEN")

	if requestAuthToken == "" {
		return errors.New("token is missing in the request")
	}

	if requestAuthToken != authToken {
		return errors.New("unauthorised request")
	}

	return nil
}
