package api

import (
	"net/http"
)

type Config struct{}

func (app *Config) BrokerHandler(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
