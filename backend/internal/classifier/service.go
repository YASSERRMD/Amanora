package classifier

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
	"github.com/YASSERRMD/Amanora/backend/internal/pii"
)

type AuditSink interface {
	RecordClassificationAudit(ctx context.Context, job Job, action string) error
}

type Service struct {
	detectors []pii.Detector
	audit     AuditSink
	mu        sync.RWMutex
	jobs      map[string]Job
	findings  []Finding
	nextID    int
}

func NewService(detectors []pii.Detector, audit AuditSink) *Service {
	return &Service{
		detectors: detectors,
		audit:     audit,
		jobs:      map[string]Job{},
		findings:  []Finding{},
		nextID:    1,
	}
}

func NewDefaultService() *Service {
	return NewService(pii.DefaultDetectors(), &MemoryAuditSink{})
}

func (s *Service) CreateJob(assets []datasource.AssetMetadata) Job {
	s.mu.Lock()
	defer s.mu.Unlock()

	job := Job{
		ID:        fmt.Sprintf("class_%06d", s.nextID),
		Status:    JobStatusPending,
		Assets:    assets,
		CreatedAt: time.Now().UTC(),
	}
	s.nextID++
	s.jobs[job.ID] = job
	return job
}

func (s *Service) RunJob(ctx context.Context, id string) (Job, error) {
	s.mu.RLock()
	job, ok := s.jobs[id]
	s.mu.RUnlock()
	if !ok {
		return Job{}, fmt.Errorf("classification job not found")
	}

	started := time.Now().UTC()
	job.Status = JobStatusRunning
	job.StartedAt = started
	s.save(job)

	findings := s.ClassifyAssets(job.Assets)
	completed := time.Now().UTC()
	job.Status = JobStatusCompleted
	job.CompletedAt = completed
	job.Findings = findings
	s.save(job)
	s.saveFindings(findings)
	if s.audit != nil {
		_ = s.audit.RecordClassificationAudit(ctx, job, "classification.job.completed")
	}
	return job, nil
}

func (s *Service) ClassifyAssets(assets []datasource.AssetMetadata) []Finding {
	findings := []Finding{}
	for _, asset := range assets {
		for _, field := range asset.Fields {
			samples := []string{field.Name, field.DataType}
			for _, detector := range s.detectors {
				for _, match := range detector.Detect(field, samples) {
					findings = append(findings, Finding{
						AssetName:  asset.Name,
						FieldName:  field.Name,
						PIIType:    match.Type,
						Detector:   match.Detector,
						Confidence: scoreConfidence(match.Confidence, field),
						Evidence:   strings.Join(match.Evidence, "; "),
						DetectedAt: time.Now().UTC(),
					})
				}
			}
		}
	}
	return findings
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

func (s *Service) ListFindings() []Finding {
	s.mu.RLock()
	defer s.mu.RUnlock()

	findings := make([]Finding, len(s.findings))
	copy(findings, s.findings)
	return findings
}

func (s *Service) save(job Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[job.ID] = job
}

func (s *Service) saveFindings(findings []Finding) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.findings = append(s.findings, findings...)
}

func scoreConfidence(base float64, field datasource.FieldMetadata) float64 {
	score := base
	if field.DataType == "text" {
		score += 0.02
	}
	if score > 1 {
		return 1
	}
	return score
}

type MemoryAuditSink struct {
	mu      sync.Mutex
	Actions []string
}

func (a *MemoryAuditSink) RecordClassificationAudit(ctx context.Context, job Job, action string) error {
	_ = ctx
	_ = job
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Actions = append(a.Actions, action)
	return nil
}
