package server

import (
	"context"
	"net/http"
	"time"

	"github.com/Jira-Analyzer/backend-services/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewServer(config *config.ServerConfig, router *mux.Router) *Server {
	var (
		readTimeout     = 5 * time.Second
		writeTimeout    = 5 * time.Second
		addr            = ":8000"
		shutdownTimeout = 3 * time.Second
	)

	if config.Host != nil {
		addr = *config.Host
	}

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	serv := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: shutdownTimeout,
	}

	return serv
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) GetNotify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
