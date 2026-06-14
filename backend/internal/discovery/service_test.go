package discovery

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

func TestDiscoveryJobLifecycleStoresResultAndAudit(t *testing.T) {
	dir := t.TempDir()
	csvPath := filepath.Join(dir, "customers.csv")
	if err := os.WriteFile(csvPath, []byte("email,age\nperson@example.com,42\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	registry, err := datasource.NewRegistry(datasource.NewCSVConnector())
	if err != nil {
		t.Fatal(err)
	}
	dataSources := datasource.NewService(registry)
	source, err := dataSources.Register(datasource.DataSource{
		Name: "customers",
		Type: datasource.SourceTypeCSV,
		Options: map[string]string{
			"path": csvPath,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	events := &MemoryEventPublisher{}
	audit := &MemoryAuditSink{}
	service := NewService(dataSources, events, audit)

	job, err := service.CreateJob(context.Background(), source.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != JobStatusPending {
		t.Fatalf("expected pending job, got %s", job.Status)
	}

	job, err = service.RunJob(context.Background(), job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != JobStatusCompleted {
		t.Fatalf("expected completed job, got %s", job.Status)
	}

	result, ok := service.GetResult(job.ID)
	if !ok {
		t.Fatal("expected stored discovery result")
	}
	if len(result.Assets) != 1 {
		t.Fatalf("expected 1 discovered asset, got %d", len(result.Assets))
	}
	if len(events.Events) != 3 {
		t.Fatalf("expected 3 events, got %d", len(events.Events))
	}
	if len(audit.Actions) != 3 {
		t.Fatalf("expected 3 audit actions, got %d", len(audit.Actions))
	}
}
