package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Host             *string        `yaml:"host"`
	ResourceTimeout  *time.Duration `yaml:"resource-timeout"`
	AnalyticsTimeout *time.Duration `yaml:"analytics-timeout"`
	ReadTimeout      *time.Duration `yaml:"read-timeout"`
	WriteTimeout     *time.Duration `yaml:"write-timeout"`
}

type Server struct {
	server *http.Server
	notify chan error
}

func NewServer(config *ServerConfig, router *mux.Router, notify chan error) *Server {
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
