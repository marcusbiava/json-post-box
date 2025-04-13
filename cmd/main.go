package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/marcusbiava/json-post-box/internal/adapter/http/handler"
	"github.com/marcusbiava/json-post-box/internal/adapter/repository/memory"
	"github.com/marcusbiava/json-post-box/internal/infrastructure/server"
	"github.com/marcusbiava/json-post-box/internal/usecase"
)

func main() {
	port := "8030"
	// Create context that listens for interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize dependencies
	jsonRepo := memory.NewJSONRepository()
	jsonUsecase := usecase.NewJSONUsecase(jsonRepo)
	jsonHandler := handler.NewJSONHandler(jsonUsecase)

	// Create and start server
	srv := server.NewEchoServer(jsonHandler, port)

	if err := srv.Start(ctx); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
