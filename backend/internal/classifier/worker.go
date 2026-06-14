package classifier

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

	logger.Info("classifier worker ready")
	for {
		select {
		case <-ctx.Done():
			logger.Info("classifier worker stopped")
			return ctx.Err()
		case <-ticker.C:
			logger.Debug("classifier worker heartbeat")
		}
	}
}
