package httper

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config is type for server setup.
type ServerCfg struct {
	Url             string        `yaml:"url"             env:"SERVER_URL"              env-default:":80"`
	ReadTimeout     time.Duration `yaml:"readTimeout"     env:"SERVER_READ_TIMEOUT"     env-default:"5s"`
	WriteTimeout    time.Duration `yaml:"writeTimeout"    env:"SERVER_WRITE_TIMEOUT"    env-default:"5s"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" env:"SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
}

// Server represents an HTTP server with basic functionalities.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// NewServer creates a new instance of Server with the given configuration and handler.
func NewServer(cfg *ServerCfg, handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Addr:         cfg.Url,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.ShutdownTimeout,
	}

	return s
}

// Start begins listening for incoming HTTP requests in a separate goroutine.
func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify returns a read-only channel to receive notifications about server errors.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown gracefully shuts down the server, waiting for ongoing requests to finish.
func (s *Server) Shutdown(log *slog.Logger) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("signal: " + s.String())
	case err := <-s.Notify():
		log.Error("httpServer.Notify:" + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
