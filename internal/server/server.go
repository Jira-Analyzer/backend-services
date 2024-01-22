package server

import (
	"net/http"

	"github.com/Jira-Analyzer/backend-services/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	server *http.Server
	notify chan error
}

func NewServer(config *config.ServerConfig, router *mux.Router, notify chan error) *Server {
	httpServer := &http.Server{
		Addr:         *config.Host,
		Handler:      router,
		ReadTimeout:  *config.ReadTimeout,
		WriteTimeout: *config.WriteTimeout,
	}

	serv := &Server{
		server: httpServer,
		notify: notify,
	}

	return serv
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
	}()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
