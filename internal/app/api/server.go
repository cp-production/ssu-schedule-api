package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

func newServer(logger *logrus.Logger, store *store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.logRequest)
 	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	s.router.HandleFunc("/api/v1.0/departments", s.handleDepartments())
	s.router.HandleFunc("/api/v1.0/{ed_form}/{dep_url}/groups", s.handleGroups())
	s.router.HandleFunc("/api/v1.0/{ed_form}/{dep_url}/{group_num}", s.handleStudentsSchedule())
	s.router.HandleFunc("/api/v1.0/{ed_form}/{dep_url}/{group_num}/subgroups", s.handleSubgroups())
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("Started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"Completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
