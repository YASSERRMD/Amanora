package datasource

import (
	"context"
	"testing"
)

func TestPlaceholderConnectorReturnsExplicitResult(t *testing.T) {
	connector := NewPlaceholderConnector(SourceTypeOracle, "oracle")
	result, err := connector.TestConnection(context.Background(), ConnectionConfig{Type: SourceTypeOracle})
	if err != nil {
		t.Fatalf("test connection: %v", err)
	}
	if result.OK {
		t.Fatal("expected placeholder connector not to report ok")
	}
	if result.Message == "" {
		t.Fatal("expected placeholder message")
	}
}
