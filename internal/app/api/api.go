package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cp-production/ssu-schedule-api/internal/app/parser"

	_ "github.com/cp-production/ssu-schedule-api/internal/app/api/model"
	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
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
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	s.router.HandleFunc("/api/v1.0/departments", s.handleDepartments())
	s.router.HandleFunc("/api/v1.0/{ed_form}/{dep_url}/groups", s.handleGroups())
	s.router.HandleFunc("/api/v1.0/{ed_form}/{dep_url}/{group_num}", s.handleStudentsSchedule())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	start := time.Now()
	err := parser.ParseAll(s.store)
	if err != nil {
		return err
	}
	s.Logger.Info("Parsed SSU Schedule in ", time.Since(start))

	return nil
}

// @Summary get a list of departments
// @ID get-departments-list
// @Tags departments
// @Description Retrieves SSU departments' list
// @Produce json
// @Success 200 {array} model.Departments
// @Router /departments [get]
func (s *Server) handleDepartments() http.HandlerFunc {
	d, _ := s.store.Departments().SelectAll()

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, *d)
	}
}

// @Summary get a list of groups of a certain department
// @Tags groups
// @Description Retrieves groups' list based on department and education form
// @ID get-groups-list
// @Param education_form path string true "Education form, e.g. `do`"
// @Param department path string true "Department URL, e.g. `knt` for CSIT department"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Groups
// @Router /{education_form}/{department}/groups [get]
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

// @Summary get the schedule of students for a particular group
// @Tags schedule
// @Description Retrieves the schedule based on department, education form and group number
// @ID get-students-schedule
// @Param education_form path string true "Education form, e.g. `do`"
// @Param department path string true "Department URL, e.g. `knt` for CSIT department"
// @Param group_number path string true "Group number, e.g. `351`"
// @Accept  json
// @Produce  json
// @Success 200 {array} model.StudentsLesson
// @Router /{education_form}/{department}/{group_number} [get]
func (s *Server) handleStudentsSchedule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		educationForm := vars["ed_form"]
		departmentUrl := vars["dep_url"]
		groupNum := vars["group_num"]
		l, err := s.store.StudentsSchedule().Select(educationForm, departmentUrl, groupNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if l == nil {
			http.Error(w, "Group not found", http.StatusNotFound)
			return
		}
		s.respond(w, http.StatusOK, *l)
	}
}

func (s *Server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.Logger.Info("ERROR")
		}
	}
}
