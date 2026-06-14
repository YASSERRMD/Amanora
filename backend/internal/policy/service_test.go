package policy

import (
	"context"
	"testing"
)

func TestPolicyEvaluation(t *testing.T) {
	service := NewService()
	_, err := service.AddPolicy(context.Background(), Document{
		Name:       "PII must have owner",
		Severity:   "high",
		Conditions: map[string]string{"classification": "pii", "owner": "present"},
		Effect:     "warn",
	})
	if err != nil {
		t.Fatal(err)
	}
	decision, err := service.Evaluate(context.Background(), "PII must have owner", map[string]string{
		"classification": "pii",
		"owner":          "missing",
	})
	if err != nil {
		t.Fatal(err)
	}
	if decision.Compliant {
		t.Fatal("expected non-compliant decision")
	}
	if len(service.EvaluateAll(context.Background(), map[string]string{"classification": "pii"})) != 1 {
		t.Fatal("expected evaluate-all decision")
	}
}
