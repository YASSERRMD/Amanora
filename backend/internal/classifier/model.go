package classifier

import (
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)

type Job struct {
	ID          string                     `json:"id"`
	Status      JobStatus                  `json:"status"`
	Assets      []datasource.AssetMetadata `json:"assets,omitempty"`
	Findings    []Finding                  `json:"findings,omitempty"`
	Error       string                     `json:"error,omitempty"`
	CreatedAt   time.Time                  `json:"createdAt"`
	StartedAt   time.Time                  `json:"startedAt,omitempty"`
	CompletedAt time.Time                  `json:"completedAt,omitempty"`
}

type Finding struct {
	AssetName  string    `json:"assetName"`
	FieldName  string    `json:"fieldName"`
	PIIType    string    `json:"piiType"`
	Detector   string    `json:"detector"`
	Confidence float64   `json:"confidence"`
	Evidence   string    `json:"evidence,omitempty"`
	DetectedAt time.Time `json:"detectedAt"`
}

type Classifier interface {
	Classify(asset datasource.AssetMetadata) []Finding
}
