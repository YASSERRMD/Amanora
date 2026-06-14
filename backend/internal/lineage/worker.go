package lineage

import (
	"context"

	"github.com/YASSERRMD/Amanora/backend/internal/graph"
)

type SyncWorker struct {
	service *Service
}

func NewSyncWorker(service *Service) SyncWorker {
	return SyncWorker{service: service}
}

func (w SyncWorker) SyncEdges(ctx context.Context, edges []graph.Edge) error {
	for _, edge := range edges {
		if _, err := w.service.UpsertEdge(ctx, edge); err != nil {
			return err
		}
	}
	return nil
}
