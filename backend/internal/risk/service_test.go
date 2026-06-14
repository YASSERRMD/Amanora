package risk

import "testing"

func TestRiskScoreCalculation(t *testing.T) {
	score := NewService().Calculate(AssetSignal{
		AssetID:          "asset_1",
		PiiFindings:      3,
		HasOwner:         false,
		HasRetention:     false,
		ComplianceFailed: 1,
	})
	if score.Level != "high" && score.Level != "critical" {
		t.Fatalf("expected elevated risk, got %s", score.Level)
	}
	if score.Score <= 0 {
		t.Fatal("expected positive score")
	}
}
