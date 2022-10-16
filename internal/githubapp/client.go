package githubapp

import (
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v45/github"
	"github.com/ryanrolds/gh_bot/internal/config"
)

func NewClient(cfg *config.Config) (*github.Client, error) {
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, cfg.GhAppID, cfg.GhInstllationID, ".github-app-key.pem")
	if err != nil {
		return nil, err
	}

	// Use installation transport with client.
	client := github.NewClient(&http.Client{Transport: itr})

	return client, nil
}
