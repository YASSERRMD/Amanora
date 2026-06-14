package discovery

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

func TestDiscoveryJobLifecycleStoresResultsAndEvents(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "customers.csv")
	if err := os.WriteFile(path, []byte("email,age\nada@example.com,42\n"), 0o600); err != nil {
		t.Fatalf("write csv: %v", err)
	}

	registry, err := datasource.DefaultRegistry()
	if err != nil {
		t.Fatalf("default registry: %v", err)
	}
	dataSources := datasource.NewService(registry)
	source, err := dataSources.Register(datasource.DataSource{
		Name:    "customers",
		Type:    datasource.SourceTypeCSV,
		Options: map[string]string{"path": path},
	})
	if err != nil {
		t.Fatalf("register source: %v", err)
	}

	publisher := &MemoryEventPublisher{}
	audit := &MemoryAuditSink{}
	service := NewService(dataSources, publisher, audit)

	job, err := service.CreateJob(context.Background(), source.ID)
	if err != nil {
		t.Fatalf("create job: %v", err)
	}
	if job.Status != JobStatusPending {
		t.Fatalf("expected pending job, got %s", job.Status)
	}

	job, err = service.RunJob(context.Background(), job.ID)
	if err != nil {
		t.Fatalf("run job: %v", err)
	}
	if job.Status != JobStatusCompleted {
		t.Fatalf("expected completed job, got %s", job.Status)
	}

	result, ok := service.GetResult(job.ID)
	if !ok {
		t.Fatal("expected discovery result")
	}
	if len(result.Assets) != 1 {
		t.Fatalf("expected one discovered asset, got %d", len(result.Assets))
	}
	if len(publisher.Snapshot()) != 3 {
		t.Fatalf("expected 3 discovery events, got %d", len(publisher.Snapshot()))
	}
	if len(audit.Snapshot()) != 3 {
		t.Fatalf("expected 3 audit actions, got %d", len(audit.Snapshot()))
	}
}

func TestDiscoveryMetadataSummaries(t *testing.T) {
	assets := []datasource.AssetMetadata{{
		Name:       "customers",
		SchemaName: "public",
		Fields: []datasource.FieldMetadata{
			{Name: "email", DataType: "text", Nullable: false},
			{Name: "age", DataType: "integer", Nullable: true},
		},
	}}

	if len(NewSchemaDiscoveryService().DiscoverSchemas(assets)) != 1 {
		t.Fatal("expected one schema summary")
	}
	if len(NewTableDiscoveryService().DiscoverTables(assets)) != 1 {
		t.Fatal("expected one table summary")
	}
	if len(NewFieldDiscoveryService().DiscoverFields(assets)) != 2 {
		t.Fatal("expected two field summaries")
	}
	if len(NewSamplingService().BuildSamples(assets)) != 2 {
		t.Fatal("expected two sample sets")
	}
}
