package datasource

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestCSVConnectorDiscoversFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "customers.csv")
	if err := os.WriteFile(path, []byte("email,age,active\nada@example.com,42,true\n"), 0o600); err != nil {
		t.Fatalf("write csv: %v", err)
	}

	connector := NewCSVConnector()
	assets, err := connector.DiscoverAssets(context.Background(), ConnectionConfig{
		Name:    "sample",
		Type:    SourceTypeCSV,
		Options: map[string]string{"path": path},
	})
	if err != nil {
		t.Fatalf("discover assets: %v", err)
	}

	if len(assets) != 1 {
		t.Fatalf("expected 1 asset, got %d", len(assets))
	}
	if len(assets[0].Fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(assets[0].Fields))
	}
	if assets[0].Fields[1].DataType != "integer" {
		t.Fatalf("expected age to infer integer, got %s", assets[0].Fields[1].DataType)
	}
}
