package graph

import (
	"context"
	"sync"
)

type Node struct {
	ID         string            `json:"id"`
	AssetID    string            `json:"assetId,omitempty"`
	Type       string            `json:"type"`
	Label      string            `json:"label"`
	Properties map[string]string `json:"properties,omitempty"`
}

type Edge struct {
	ID       string            `json:"id"`
	FromID   string            `json:"fromId"`
	ToID     string            `json:"toId"`
	Type     string            `json:"type"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type Repository interface {
	UpsertNode(ctx context.Context, node Node) (Node, error)
	UpsertEdge(ctx context.Context, edge Edge) (Edge, error)
	Upstream(ctx context.Context, nodeID string) ([]Edge, error)
	Downstream(ctx context.Context, nodeID string) ([]Edge, error)
}

type MemoryRepository struct {
	mu    sync.RWMutex
	nodes map[string]Node
	edges map[string]Edge
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{nodes: map[string]Node{}, edges: map[string]Edge{}}
}

func (r *MemoryRepository) UpsertNode(ctx context.Context, node Node) (Node, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nodes[node.ID] = node
	return node, nil
}

func (r *MemoryRepository) UpsertEdge(ctx context.Context, edge Edge) (Edge, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()
	r.edges[edge.ID] = edge
	return edge, nil
}

func (r *MemoryRepository) Upstream(ctx context.Context, nodeID string) ([]Edge, error) {
	_ = ctx
	r.mu.RLock()
	defer r.mu.RUnlock()
	edges := []Edge{}
	for _, edge := range r.edges {
		if edge.ToID == nodeID {
			edges = append(edges, edge)
		}
	}
	return edges, nil
}

func (r *MemoryRepository) Downstream(ctx context.Context, nodeID string) ([]Edge, error) {
	_ = ctx
	r.mu.RLock()
	defer r.mu.RUnlock()
	edges := []Edge{}
	for _, edge := range r.edges {
		if edge.FromID == nodeID {
			edges = append(edges, edge)
		}
	}
	return edges, nil
}
