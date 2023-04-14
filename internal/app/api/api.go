package api

import (
	"io"
	"net/http"

	"github.com/cp-production/ssu-schedule-api/internal/app/store"
	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	router *mux.Router
	store *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc(("/hello"), s.handlerHello())
}

func (s *Server) configureStore() error{
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *Server) handlerHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
