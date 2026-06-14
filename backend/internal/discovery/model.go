package discovery

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
	ID           string    `json:"id"`
	DataSourceID string    `json:"dataSourceId"`
	Status       JobStatus `json:"status"`
	Error        string    `json:"error,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	StartedAt    time.Time `json:"startedAt,omitempty"`
	CompletedAt  time.Time `json:"completedAt,omitempty"`
}

type Result struct {
	JobID        string                     `json:"jobId"`
	DataSourceID string                     `json:"dataSourceId"`
	Assets       []datasource.AssetMetadata `json:"assets"`
	DiscoveredAt time.Time                  `json:"discoveredAt"`
}
