package server

import (
	"context"
	"github.com/marcusbiava/json-post-box/internal/adapter/http/handler"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	e      *echo.Echo
	port   string
	server *http.Server
}

func NewEchoServer(h *handler.JSONHandler, port string) *EchoServer {
	e := echo.New()
	SetupRoutes(e, h)

	return &EchoServer{
		e:    e,
		port: ":" + port,
		server: &http.Server{
			Addr:    ":" + port,
			Handler: e,
		},
	}
}

func (s *EchoServer) Start(ctx context.Context) error {
	// Channel to listen for errors
	errChan := make(chan error, 1)

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", s.port)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Wait for interrupt signal or context cancellation
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Server shutdown by context")
	case err := <-errChan:
		log.Printf("Server error: %v", err)
		return err
	case <-quit:
		log.Println("Server shutdown by signal")
	}

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := s.server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
