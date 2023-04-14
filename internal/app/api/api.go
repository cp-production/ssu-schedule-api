package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc(("/hello"), s.handlerHello())
}

func (s *Server) handlerHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
