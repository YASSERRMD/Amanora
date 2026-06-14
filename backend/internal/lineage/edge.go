package lineage

import (
	"context"

	"github.com/YASSERRMD/Amanora/backend/internal/graph"
)

type EdgeService struct {
	service *Service
}

func NewEdgeService(service *Service) EdgeService {
	return EdgeService{service: service}
}

func (s EdgeService) CreateEdge(ctx context.Context, edge graph.Edge) (graph.Edge, error) {
	return s.service.UpsertEdge(ctx, edge)
}
