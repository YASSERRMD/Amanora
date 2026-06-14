package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/YASSERRMD/Amanora/backend/internal/discovery"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := (discovery.Worker{Logger: logger}).Run(ctx); err != nil && err != context.Canceled {
		logger.Error("discovery worker failed", "error", err)
		os.Exit(1)
	}
}
