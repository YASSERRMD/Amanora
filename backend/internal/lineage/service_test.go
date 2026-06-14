package lineage

import (
	"context"
	"testing"

	"github.com/YASSERRMD/Amanora/backend/internal/graph"
)

func TestLineageTraversal(t *testing.T) {
	service := NewService(graph.NewMemoryRepository())
	_, err := service.UpsertEdge(context.Background(), graph.Edge{FromID: "raw_customers", ToID: "customers", Type: "transforms"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = service.UpsertEdge(context.Background(), graph.Edge{FromID: "customers", ToID: "customer_report", Type: "produces"})
	if err != nil {
		t.Fatal(err)
	}
	upstream, err := service.Upstream(context.Background(), "customers")
	if err != nil {
		t.Fatal(err)
	}
	downstream, err := service.Downstream(context.Background(), "customers")
	if err != nil {
		t.Fatal(err)
	}
	if len(upstream) != 1 || len(downstream) != 1 {
		t.Fatalf("expected one upstream and downstream edge, got %d/%d", len(upstream), len(downstream))
	}
}
