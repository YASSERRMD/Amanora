package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/config"
	httpapi "github.com/YASSERRMD/Amanora/backend/internal/http"
)

func main() {
	cfg := config.Load()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewRouter(cfg),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errs := make(chan error, 1)
	go func() {
		logger.Info("starting api", "addr", cfg.HTTPAddr, "environment", cfg.Environment)
		errs <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errs:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("api failed", "error", err)
			os.Exit(1)
		}
	case sig := <-shutdown:
		logger.Info("shutting down api", "signal", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("api shutdown failed", "error", err)
			os.Exit(1)
		}
	}
}
