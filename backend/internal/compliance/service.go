package compliance

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Framework struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Jurisdiction string `json:"jurisdiction,omitempty"`
}

type Check struct {
	Key      string `json:"key"`
	Title    string `json:"title"`
	Severity string `json:"severity"`
}

type Result struct {
	Framework string    `json:"framework"`
	CheckKey  string    `json:"checkKey"`
	Status    string    `json:"status"`
	Severity  string    `json:"severity"`
	Evidence  string    `json:"evidence,omitempty"`
	RunAt     time.Time `json:"runAt"`
}

type Service struct {
	mu         sync.RWMutex
	frameworks map[string]Framework
	checks     []Check
	results    []Result
}

func NewService() *Service {
	return &Service{
		frameworks: map[string]Framework{
			"GDPR_STYLE": {ID: "GDPR_STYLE", Name: "GDPR-style Privacy Checks", Code: "GDPR_STYLE", Jurisdiction: "EU"},
			"UAE_PDPL":   {ID: "UAE_PDPL", Name: "UAE PDPL-style Privacy Checks", Code: "UAE_PDPL", Jurisdiction: "UAE"},
			"INDIA_DPDP": {ID: "INDIA_DPDP", Name: "India DPDP-style Privacy Checks", Code: "INDIA_DPDP", Jurisdiction: "India"},
			"INTERNAL":   {ID: "INTERNAL", Name: "Internal Enterprise Policy Checks", Code: "INTERNAL", Jurisdiction: "Global"},
		},
		checks: []Check{
			{Key: "privacy.pii_owner", Title: "PII assets have an owner", Severity: "high"},
			{Key: "ownership.steward", Title: "Assets have assigned steward", Severity: "medium"},
			{Key: "retention.assigned", Title: "High-risk assets have retention policy", Severity: "high"},
		},
	}
}

func (s *Service) Frameworks() []Framework {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []Framework{}
	for _, framework := range s.frameworks {
		out = append(out, framework)
	}
	return out
}

func (s *Service) Run(ctx context.Context, facts map[string]string) []Result {
	_ = ctx
	results := []Result{}
	for code := range s.frameworks {
		for _, check := range s.checks {
			status := "pass"
			if facts[check.Key] == "fail" {
				status = "fail"
			}
			results = append(results, Result{
				Framework: code,
				CheckKey:  check.Key,
				Status:    status,
				Severity:  check.Severity,
				Evidence:  fmt.Sprintf("fact=%s", facts[check.Key]),
				RunAt:     time.Now().UTC(),
			})
		}
	}
	s.mu.Lock()
	s.results = append(s.results, results...)
	s.mu.Unlock()
	return results
}

func (s *Service) Results() []Result {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Result, len(s.results))
	copy(out, s.results)
	return out
}
