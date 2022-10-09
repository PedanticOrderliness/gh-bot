package api

import (
	"net/http"

	"github.com/ryanrolds/gh_bot/internal/config"
	"github.com/sirupsen/logrus"
)

type API struct {
	config *config.Config
}

func NewAPI(cfg *config.Config) *API {
	return &API{
		config: cfg,
	}
}

func writeResponse(w http.ResponseWriter, status int, reason string) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(reason))
	if err != nil {
		logrus.WithError(err).Warn("problem writing response")
	}
}
