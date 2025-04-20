package main

import (
	"context"
	"fmt"
	"github/milankatira/students-api-go/internal/config"
	"github/milankatira/students-api-go/internal/http/handler/student"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", student.New())

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server starting on %s\n", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("fail to start server")
		}
	}()

	<-done
	slog.Info("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		log.Fatal("server shutdown failed")
	}

	slog.Info("server exited properly")
}
