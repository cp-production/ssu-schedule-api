package api

import (
	"net/http"
	"time"

	// "github.com/cp-production/ssu-schedule-api/internal/app/parser"

	_ "github.com/cp-production/ssu-schedule-api/internal/app/api/model"
	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

func Start(config *Config) error {
	logger := logrus.New()
	if err := configureLogger(config, logger); err != nil {
		return err
	}
	store, err := configureStore(config, logger)
	if err != nil {
		return err
	}
	srv := newServer(logger, store)

	logger.Info("Server is listening")
	return http.ListenAndServe(config.BindAddr, srv)
}

func configureLogger(config *Config, logger *logrus.Logger) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	logger.SetLevel(level)
	return nil
}

func configureStore(config *Config, logger *logrus.Logger) (*store.Store, error) {
	st := store.New(config.Store)
	if err := st.Open(); err != nil {
		return nil, err
	}

	start := time.Now()
	// err := parser.ParseAll(st)
	// if err != nil {
	// 	return nil, err
	// }
	logger.Info("Parsed SSU info in ", time.Since(start))

	return st, nil 
}