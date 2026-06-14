package lineage

import (
	"context"
	"fmt"

	"github.com/YASSERRMD/Amanora/backend/internal/graph"
)

type Service struct {
	repo graph.Repository
}

func NewService(repo graph.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) UpsertNode(ctx context.Context, node graph.Node) (graph.Node, error) {
	if node.ID == "" {
		return graph.Node{}, fmt.Errorf("node id is required")
	}
	return s.repo.UpsertNode(ctx, node)
}

func (s *Service) UpsertEdge(ctx context.Context, edge graph.Edge) (graph.Edge, error) {
	if edge.ID == "" {
		edge.ID = edge.FromID + "_to_" + edge.ToID + "_" + edge.Type
	}
	return s.repo.UpsertEdge(ctx, edge)
}

func (s *Service) AssetLineage(ctx context.Context, assetID string) (map[string]any, error) {
	upstream, err := s.repo.Upstream(ctx, assetID)
	if err != nil {
		return nil, err
	}
	downstream, err := s.repo.Downstream(ctx, assetID)
	if err != nil {
		return nil, err
	}
	return map[string]any{"assetId": assetID, "upstream": upstream, "downstream": downstream}, nil
}

func (s *Service) Upstream(ctx context.Context, assetID string) ([]graph.Edge, error) {
	return s.repo.Upstream(ctx, assetID)
}

func (s *Service) Downstream(ctx context.Context, assetID string) ([]graph.Edge, error) {
	return s.repo.Downstream(ctx, assetID)
}
