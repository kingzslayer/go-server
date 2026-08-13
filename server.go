package server

import (
	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	http            *http.Server
	shutdownTimeout time.Duration
}

type ServerConfig struct {
	Addr            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func NewServer(cfg ServerConfig, handler http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:         cfg.Addr,
			Handler:      handler,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errc := make(chan error, 1)

	go func() {
		if err := s.http.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			errc <- err
		}
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		s.shutdownTimeout,
	)
	defer cancel()

	return s.http.Shutdown(shutdownCtx)
}
