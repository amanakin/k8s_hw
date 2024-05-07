package main

import (
	"context"
	"github.com/amanakin/k8s_hw/internal/handler"
	"github.com/amanakin/k8s_hw/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := service.NewInMemoryService()
	h := handler.NewHandler(s)
	server := &http.Server{
		Addr:    "0.0.0.0:12121",
		Handler: newHandler(h),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	go func() {
		log.Println("Starting server on http://localhost:12121")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server exited properly")
}

func newHandler(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/time", h.HandleTime)
	r.Get("/statistics", h.HandleStatistics)

	return r
}
