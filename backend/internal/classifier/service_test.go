package classifier

import (
	"context"
	"testing"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
	"github.com/YASSERRMD/Amanora/backend/internal/pii"
)

func TestServiceClassifiesEmailAndPhone(t *testing.T) {
	audit := &MemoryAuditSink{}
	service := NewService(pii.DefaultDetectors(), audit)
	job := service.CreateJob([]datasource.AssetMetadata{{
		Name: "customers",
		Fields: []datasource.FieldMetadata{
			{Name: "email", DataType: "text"},
			{Name: "phone_number", DataType: "text"},
		},
	}})

	job, err := service.RunJob(context.Background(), job.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.Status != JobStatusCompleted {
		t.Fatalf("expected completed job, got %s", job.Status)
	}
	if len(job.Findings) < 2 {
		t.Fatalf("expected at least 2 findings, got %d", len(job.Findings))
	}
	if len(service.ListFindings()) != len(job.Findings) {
		t.Fatal("expected findings to be stored")
	}
	if len(audit.Actions) != 1 {
		t.Fatalf("expected audit action, got %d", len(audit.Actions))
	}
}
