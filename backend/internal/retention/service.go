package retention

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Policy struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	RetentionDays int    `json:"retentionDays"`
	Action        string `json:"action"`
	Rule          string `json:"rule,omitempty"`
}

type Assignment struct {
	AssetID    string    `json:"assetId"`
	PolicyID   string    `json:"policyId"`
	AssignedAt time.Time `json:"assignedAt"`
}

type Finding struct {
	AssetID     string    `json:"assetId"`
	PolicyID    string    `json:"policyId"`
	ExpiredAt   time.Time `json:"expiredAt"`
	Violation   bool      `json:"violation"`
	Recommended string    `json:"recommended"`
}

type Service struct {
	mu          sync.RWMutex
	policies    map[string]Policy
	assignments map[string]Assignment
	nextID      int
}

func NewService() *Service {
	return &Service{policies: map[string]Policy{}, assignments: map[string]Assignment{}, nextID: 1}
}

func (s *Service) AddPolicy(ctx context.Context, policy Policy) (Policy, error) {
	_ = ctx
	if policy.Name == "" {
		return Policy{}, fmt.Errorf("policy name is required")
	}
	if policy.RetentionDays <= 0 {
		return Policy{}, fmt.Errorf("retention days must be positive")
	}
	if policy.ID == "" {
		policy.ID = fmt.Sprintf("ret_%06d", s.nextID)
		s.nextID++
	}
	if policy.Action == "" {
		policy.Action = "review"
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.policies[policy.ID] = policy
	return policy, nil
}

func (s *Service) ListPolicies() []Policy {
	s.mu.RLock()
	defer s.mu.RUnlock()
	policies := []Policy{}
	for _, policy := range s.policies {
		policies = append(policies, policy)
	}
	return policies
}

func (s *Service) Assign(assetID string, policyID string) (Assignment, error) {
	s.mu.RLock()
	_, ok := s.policies[policyID]
	s.mu.RUnlock()
	if !ok {
		return Assignment{}, fmt.Errorf("retention policy not found")
	}
	assignment := Assignment{AssetID: assetID, PolicyID: policyID, AssignedAt: time.Now().UTC()}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.assignments[assetID] = assignment
	return assignment, nil
}

func (s *Service) Evaluate(assetID string, createdAt time.Time) (Finding, error) {
	s.mu.RLock()
	assignment, ok := s.assignments[assetID]
	policy := s.policies[assignment.PolicyID]
	s.mu.RUnlock()
	if !ok {
		return Finding{}, fmt.Errorf("asset has no retention assignment")
	}
	expiredAt := CalculateExpiry(createdAt, policy.RetentionDays)
	violation := time.Now().UTC().After(expiredAt)
	return Finding{AssetID: assetID, PolicyID: policy.ID, ExpiredAt: expiredAt, Violation: violation, Recommended: policy.Action}, nil
}

func (s *Service) Simulate(assetID string, createdAt time.Time, at time.Time) (Finding, error) {
	finding, err := s.Evaluate(assetID, createdAt)
	if err != nil {
		return Finding{}, err
	}
	finding.Violation = at.After(finding.ExpiredAt)
	return finding, nil
}

func CalculateExpiry(createdAt time.Time, retentionDays int) time.Time {
	return createdAt.AddDate(0, 0, retentionDays)
}

func IsViolation(expiredAt time.Time, at time.Time) bool {
	return at.After(expiredAt)
}
