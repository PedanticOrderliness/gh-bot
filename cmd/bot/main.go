package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ryanrolds/gh_bot/internal/api"
	"github.com/ryanrolds/gh_bot/internal/config"
	"github.com/ryanrolds/gh_bot/internal/githubapp"
	"github.com/sirupsen/logrus"
)

const (
	configFilename = "bot.yaml"
)

func main() {
	cfg, err := config.GetConfig(configFilename)
	if err != nil {
		log.Fatalf("problem reading %s: %s", configFilename, err)
	}

	initLogging()

	ghAppClient, err := githubapp.NewClient(cfg)
	if err != nil {
		log.Fatalf("problem creating github client: %s", err)
	}

	a := api.NewAPI(cfg, ghAppClient)
	r := mux.NewRouter()

	// TODO add request logging middleware
	r.Use(a.MiddlewareCheckAccess)
	r.HandleFunc("/pull_request/comment", a.PullRequestComment).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.Fatal(err)
		return
	}

	logrus.Info("server gracefully shutdown")
}
