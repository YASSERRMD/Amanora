package pii

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type Detection struct {
	Type       string   `json:"type"`
	Detector   string   `json:"detector"`
	Confidence float64  `json:"confidence"`
	Evidence   []string `json:"evidence,omitempty"`
}

type Detector interface {
	Name() string
	Detect(field datasource.FieldMetadata, samples []string) []Detection
}

func DefaultDetectors() []Detector {
	return []Detector{
		NewEmailDetector(),
	}
}
