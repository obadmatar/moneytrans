package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/obadmatar/moneytrans/internal/account"
	"github.com/obadmatar/moneytrans/internal/transfer"
)

type ServerConfig  func (s *server) error

type server struct {
	logger *log.Logger
	router *httprouter.Router
	transferService *transfer.Service
	accountRepository account.Repositoy
}

// NewServer creates a new server instance and it's routes
func NewServer(configs...ServerConfig) (*server, error) {
	s := &server{}
	s.router = httprouter.New()
	s.router.NotFound = http.HandlerFunc(s.notFoundResponse)
	s.router.MethodNotAllowed = http.HandlerFunc(s.methodNotAllowed)

	for _, cfg := range configs {
		err := cfg(s)

		if err != nil {
			return nil, err
		}
	}

	// Register server routes
	s.routes()

	// Create a new logger instance to write logs to stdout
	s.logger = log.New(os.Stdout, "", log.LstdFlags)

	return s, nil
}

func WithAccountRepository(repo account.Repositoy) ServerConfig {
	return func(s *server) error {
		s.accountRepository = repo
		return nil
	}
}

func WithTransferService(service *transfer.Service) ServerConfig {
	return func(s *server) error {
		s.transferService = service
		return nil
	}
}


// Run starts our application
func (s *server) Run() {
	
	// Create a new HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.recoverPanic(s.router),
	}

	// Create a channel for interrupt signals (e.g. SIGINT, SIGTERM)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		s.logger.Printf("Starting server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			s.logger.Fatalf("listenAndServe(): %v", err)
		}
	}()

	// Wait for an interrupt signal
	<-interrupt

	// Create a context for shutdown operation with a 5s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	s.logger.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Fatalf("Shutdown(): %v", err)
	}
	s.logger.Println("Server shut down gracefully")
}
