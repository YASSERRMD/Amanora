package discovery

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

type JobStatus string

const (
	JobQueued    JobStatus = "queued"
	JobRunning   JobStatus = "running"
	JobSucceeded JobStatus = "succeeded"
	JobFailed    JobStatus = "failed"
)

type Job struct {
	ID           string                     `json:"id"`
	DataSourceID string                     `json:"dataSourceId"`
	Status       JobStatus                  `json:"status"`
	Assets       []datasource.AssetMetadata `json:"assets,omitempty"`
	Error        string                     `json:"error,omitempty"`
	CreatedAt    time.Time                  `json:"createdAt"`
	StartedAt    *time.Time                 `json:"startedAt,omitempty"`
	FinishedAt   *time.Time                 `json:"finishedAt,omitempty"`
}

type EventPublisher interface {
	PublishDiscoveryEvent(ctx context.Context, event Event) error
}

type AuditSink interface {
	RecordDiscoveryAudit(ctx context.Context, job Job, action string) error
}

type Event struct {
	Type         string    `json:"type"`
	JobID        string    `json:"jobId"`
	DataSourceID string    `json:"dataSourceId"`
	OccurredAt   time.Time `json:"occurredAt"`
}

type Service struct {
	dataSources *datasource.Service
	publisher   EventPublisher
	audit       AuditSink
	mu          sync.RWMutex
	jobs        map[string]Job
	nextID      int
}

func NewService(dataSources *datasource.Service, publisher EventPublisher, audit AuditSink) *Service {
	return &Service{
		dataSources: dataSources,
		publisher:   publisher,
		audit:       audit,
		jobs:        map[string]Job{},
		nextID:      1,
	}
}

func (s *Service) CreateJob(ctx context.Context, dataSourceID string) (Job, error) {
	if _, ok := s.dataSources.Get(dataSourceID); !ok {
		return Job{}, fmt.Errorf("data source not found")
	}

	s.mu.Lock()
	job := Job{
		ID:           fmt.Sprintf("disc_%06d", s.nextID),
		DataSourceID: dataSourceID,
		Status:       JobQueued,
		CreatedAt:    time.Now().UTC(),
	}
	s.nextID++
	s.jobs[job.ID] = job
	s.mu.Unlock()

	s.publish(ctx, job, "discovery.job.created")
	return job, nil
}

func (s *Service) ListJobs() []Job {
	s.mu.RLock()
	defer s.mu.RUnlock()

	jobs := make([]Job, 0, len(s.jobs))
	for _, job := range s.jobs {
		jobs = append(jobs, job)
	}
	return jobs
}

func (s *Service) GetJob(id string) (Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	job, ok := s.jobs[id]
	return job, ok
}

func (s *Service) RunJob(ctx context.Context, id string) (Job, error) {
	s.mu.RLock()
	job, ok := s.jobs[id]
	s.mu.RUnlock()
	if !ok {
		return Job{}, fmt.Errorf("discovery job not found")
	}

	source, ok := s.dataSources.Get(job.DataSourceID)
	if !ok {
		return Job{}, fmt.Errorf("data source not found")
	}

	started := time.Now().UTC()
	job.Status = JobRunning
	job.StartedAt = &started
	s.save(job)
	s.publish(ctx, job, "discovery.job.started")

	assets, err := s.discover(ctx, source)
	finished := time.Now().UTC()
	job.FinishedAt = &finished
	if err != nil {
		job.Status = JobFailed
		job.Error = err.Error()
		s.save(job)
		s.publish(ctx, job, "discovery.job.failed")
		return job, err
	}

	job.Status = JobSucceeded
	job.Assets = assets
	s.save(job)
	s.publish(ctx, job, "discovery.job.succeeded")
	return job, nil
}

func (s *Service) discover(ctx context.Context, source datasource.DataSource) ([]datasource.AssetMetadata, error) {
	registry, err := datasource.DefaultRegistry()
	if err != nil {
		return nil, err
	}
	connector, err := registry.Get(source.Type)
	if err != nil {
		return nil, err
	}
	return connector.DiscoverAssets(ctx, datasource.ConnectionConfig{
		Name:          source.Name,
		Type:          source.Type,
		ConnectionURI: source.ConnectionURI,
		Options:       source.Options,
	})
}

func (s *Service) save(job Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

func (s *Service) publish(ctx context.Context, job Job, eventType string) {
	if s.publisher != nil {
		_ = s.publisher.PublishDiscoveryEvent(ctx, Event{
			Type:         eventType,
			JobID:        job.ID,
			DataSourceID: job.DataSourceID,
			OccurredAt:   time.Now().UTC(),
		})
	}
	if s.audit != nil {
		_ = s.audit.RecordDiscoveryAudit(ctx, job, eventType)
	}
}

type MemoryEventPublisher struct {
	mu     sync.Mutex
	Events []Event
}

func (p *MemoryEventPublisher) PublishDiscoveryEvent(ctx context.Context, event Event) error {
	_ = ctx
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Events = append(p.Events, event)
	return nil
}

type MemoryAuditSink struct {
	mu      sync.Mutex
	Actions []string
}

func (a *MemoryAuditSink) RecordDiscoveryAudit(ctx context.Context, job Job, action string) error {
	_ = ctx
	_ = job
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Actions = append(a.Actions, action)
	return nil
}
