package retention

import (
	"context"
	"testing"
	"time"
)

func TestRetentionEvaluationAndSimulation(t *testing.T) {
	service := NewService()
	policy, err := service.AddPolicy(context.Background(), Policy{Name: "PII review", RetentionDays: 30, Action: "review"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := service.Assign("asset_1", policy.ID); err != nil {
		t.Fatal(err)
	}
	createdAt := time.Now().UTC().AddDate(0, 0, -40)
	finding, err := service.Evaluate("asset_1", createdAt)
	if err != nil {
		t.Fatal(err)
	}
	if !finding.Violation {
		t.Fatal("expected retention violation")
	}
	simulation, err := service.Simulate("asset_1", createdAt, createdAt.AddDate(0, 0, 20))
	if err != nil {
		t.Fatal(err)
	}
	if simulation.Violation {
		t.Fatal("expected simulation before expiry to pass")
	}
}
