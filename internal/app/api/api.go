package api

import (
	"encoding/json"
	"net/http"
	"time"

	// "github.com/cp-production/ssu-schedule-api/internal/app/parser"

	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	Logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	s.Logger.Info("Server is listening")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.Logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc(("/api/v1.0/departments"), s.handleDepartments())
	s.router.HandleFunc(("/api/v1.0/{ed_form}/{dep_url}/groups"), s.handleGroups())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	start := time.Now()
	// err := parser.ParseAll(s.store)
	// if err != nil {
	// 	return err
	// }
	s.Logger.Info("Parsed SSU Schedule in ", time.Since(start))

	return nil
}

func (s *Server) handleDepartments() http.HandlerFunc {
	d, _ := s.store.Departments().SelectAll()

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, *d)
	}
}

func (s *Server) handleGroups() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		educationForm := vars["ed_form"]
		departmentUrl := vars["dep_url"]
		d, err := s.store.Groups().SelectByDepartments(educationForm, departmentUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if d == nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}
		s.respond(w, http.StatusOK, *d)
	}
}

func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.Logger.Info("ERROR")
		}
	}
}
