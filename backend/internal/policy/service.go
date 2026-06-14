package policy

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Document struct {
	Name       string            `json:"name"`
	Category   string            `json:"category"`
	Severity   string            `json:"severity"`
	Conditions map[string]string `json:"conditions"`
	Effect     string            `json:"effect"`
	Message    string            `json:"message"`
}

type Decision struct {
	PolicyName  string    `json:"policyName"`
	Compliant   bool      `json:"compliant"`
	Effect      string    `json:"effect"`
	Severity    string    `json:"severity"`
	Message     string    `json:"message"`
	EvaluatedAt time.Time `json:"evaluatedAt"`
}

type Service struct {
	mu       sync.RWMutex
	policies map[string]Document
}

func NewService() *Service {
	return &Service{policies: map[string]Document{}}
}

func (s *Service) AddPolicy(ctx context.Context, doc Document) (Document, error) {
	_ = ctx
	if doc.Name == "" {
		return Document{}, fmt.Errorf("policy name is required")
	}
	if doc.Effect == "" {
		doc.Effect = "warn"
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.policies[doc.Name] = doc
	return doc, nil
}

func (s *Service) ListPolicies() []Document {
	s.mu.RLock()
	defer s.mu.RUnlock()
	policies := make([]Document, 0, len(s.policies))
	for _, policy := range s.policies {
		policies = append(policies, policy)
	}
	return policies
}

func (s *Service) Evaluate(ctx context.Context, name string, facts map[string]string) (Decision, error) {
	_ = ctx
	s.mu.RLock()
	doc, ok := s.policies[name]
	s.mu.RUnlock()
	if !ok {
		return Decision{}, fmt.Errorf("policy not found")
	}
	return evaluateDocument(doc, facts), nil
}

func (s *Service) EvaluateAll(ctx context.Context, facts map[string]string) []Decision {
	_ = ctx
	decisions := []Decision{}
	for _, doc := range s.ListPolicies() {
		decisions = append(decisions, evaluateDocument(doc, facts))
	}
	return decisions
}

func evaluateDocument(doc Document, facts map[string]string) Decision {
	compliant := true
	for key, expected := range doc.Conditions {
		if facts[key] != expected {
			compliant = false
			break
		}
	}
	message := doc.Message
	if message == "" {
		message = "policy evaluated"
	}
	return Decision{
		PolicyName:  doc.Name,
		Compliant:   compliant,
		Effect:      doc.Effect,
		Severity:    doc.Severity,
		Message:     message,
		EvaluatedAt: time.Now().UTC(),
	}
}
