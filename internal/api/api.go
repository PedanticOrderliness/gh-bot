package api

import (
	"net/http"

	"github.com/google/go-github/v45/github"
	"github.com/ryanrolds/gh_bot/internal/config"
	"github.com/sirupsen/logrus"
)

type API struct {
	config      *config.Config
	ghAppClient *github.Client
}

func NewAPI(cfg *config.Config, ghAppClient *github.Client) *API {
	return &API{
		config:      cfg,
		ghAppClient: ghAppClient,
	}
}

func writeResponse(w http.ResponseWriter, status int, reason string) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(reason))
	if err != nil {
		logrus.WithError(err).Warn("problem writing response")
	}
}
