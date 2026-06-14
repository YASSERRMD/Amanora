package datasource

import "testing"

func TestRegistryRejectsDuplicateConnectors(t *testing.T) {
	_, err := NewRegistry(NewCSVConnector(), NewCSVConnector())
	if err == nil {
		t.Fatal("expected duplicate connector registration to fail")
	}
}

func TestRegistryReturnsRegisteredTypes(t *testing.T) {
	registry, err := NewRegistry(NewRESTConnector(), NewCSVConnector())
	if err != nil {
		t.Fatalf("new registry: %v", err)
	}

	types := registry.Types()
	if len(types) != 2 {
		t.Fatalf("expected 2 connector types, got %d", len(types))
	}
	if types[0] != SourceTypeCSV || types[1] != SourceTypeREST {
		t.Fatalf("unexpected connector type ordering: %#v", types)
	}
}
