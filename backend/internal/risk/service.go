package risk

type AssetSignal struct {
	AssetID          string `json:"assetId"`
	PiiFindings      int    `json:"piiFindings"`
	HasOwner         bool   `json:"hasOwner"`
	HasRetention     bool   `json:"hasRetention"`
	ComplianceFailed int    `json:"complianceFailed"`
}

type Score struct {
	AssetID string         `json:"assetId"`
	Score   int            `json:"score"`
	Level   string         `json:"level"`
	Factors map[string]int `json:"factors"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(signal AssetSignal) Score {
	score := signal.PiiFindings*15 + signal.ComplianceFailed*20
	if !signal.HasOwner {
		score += 20
	}
	if !signal.HasRetention {
		score += 15
	}
	if score > 100 {
		score = 100
	}
	level := "low"
	switch {
	case score >= 80:
		level = "critical"
	case score >= 60:
		level = "high"
	case score >= 30:
		level = "medium"
	}
	return Score{
		AssetID: signal.AssetID,
		Score:   score,
		Level:   level,
		Factors: map[string]int{
			"piiFindings":      signal.PiiFindings,
			"complianceFailed": signal.ComplianceFailed,
		},
	}
}
