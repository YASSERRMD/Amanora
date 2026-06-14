package worker

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Worker struct {
	Name   string
	Logger *slog.Logger
	Tick   time.Duration
}

func (w Worker) Run(ctx context.Context) error {
	tick := w.Tick
	if tick == 0 {
		tick = 30 * time.Second
	}

	logger := w.Logger
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	logger.Info("starting worker", "worker", w.Name, "tick", tick.String())

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logger.Info("stopping worker", "worker", w.Name)
			return ctx.Err()
		case <-ticker.C:
			logger.Debug("worker heartbeat", "worker", w.Name)
		}
	}
}

func RunUntilSignal(name string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	err := Worker{Name: name, Logger: logger}.Run(ctx)
	if err != nil && err != context.Canceled {
		logger.Error("worker stopped with error", "worker", name, "error", err)
		os.Exit(1)
	}
}
