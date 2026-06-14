package datasource

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestRegistryRejectsDuplicateConnectors(t *testing.T) {
	_, err := NewRegistry(NewCSVConnector(), NewCSVConnector())
	if err == nil {
		t.Fatal("expected duplicate connector error")
	}
}

func TestCSVConnectorDiscoversFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "customers.csv")
	if err := os.WriteFile(path, []byte("email,age,active\nperson@example.com,42,true\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	connector := NewCSVConnector()
	assets, err := connector.DiscoverAssets(context.Background(), ConnectionConfig{
		Name: "demo",
		Type: SourceTypeCSV,
		Options: map[string]string{
			"path": path,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(assets) != 1 {
		t.Fatalf("expected 1 asset, got %d", len(assets))
	}
	if got := len(assets[0].Fields); got != 3 {
		t.Fatalf("expected 3 fields, got %d", got)
	}
	if assets[0].Fields[1].DataType != "integer" {
		t.Fatalf("expected age to infer integer, got %s", assets[0].Fields[1].DataType)
	}
}

func TestServiceRegistersAndTestsPlaceholderSource(t *testing.T) {
	registry, err := NewRegistry(NewMySQLConnector())
	if err != nil {
		t.Fatal(err)
	}

	service := NewService(registry)
	source, err := service.Register(DataSource{Name: "crm", Type: SourceTypeMySQL})
	if err != nil {
		t.Fatal(err)
	}

	result, err := service.TestDataSource(context.Background(), source.ID)
	if err != nil {
		t.Fatal(err)
	}
	if result.OK {
		t.Fatal("placeholder connector should not report OK")
	}
}
