package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/tv-anagha/rest-api/internal/config"
	"github.com/tv-anagha/rest-api/internal/http/handler/student"
	"github.com/tv-anagha/rest-api/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//Database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage Initialised", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %s", err.Error())
		}
	}()

	<-done

	slog.Info("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server is shut down successfully")
}
