package discovery

import (
	"context"
	"log/slog"
	"time"
)

type Worker struct {
	Service *Service
	Logger  *slog.Logger
}

func (w Worker) Run(ctx context.Context) error {
	logger := w.Logger
	if logger == nil {
		logger = slog.Default()
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	logger.Info("discovery worker ready")
	for {
		select {
		case <-ctx.Done():
			logger.Info("discovery worker stopped")
			return ctx.Err()
		case <-ticker.C:
			logger.Debug("discovery worker heartbeat")
		}
	}
}
