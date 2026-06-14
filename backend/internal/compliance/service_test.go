package compliance

import (
	"context"
	"testing"
)

func TestComplianceRunStoresResults(t *testing.T) {
	service := NewService()
	if len(service.Frameworks()) != 4 {
		t.Fatal("expected default frameworks")
	}
	results := service.Run(context.Background(), map[string]string{"privacy.pii_owner": "fail"})
	if len(results) == 0 {
		t.Fatal("expected compliance results")
	}
	if len(service.Results()) != len(results) {
		t.Fatal("expected stored results")
	}
}
