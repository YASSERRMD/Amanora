package datasource

import (
	"context"
	"fmt"
	"sync"
)

type DataSource struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Type          SourceType        `json:"type"`
	ConnectionURI string            `json:"connectionUri,omitempty"`
	Options       map[string]string `json:"options,omitempty"`
	Status        string            `json:"status"`
}

type Service struct {
	registry *Registry
	mu       sync.RWMutex
	sources  map[string]DataSource
	nextID   int
}

func NewService(registry *Registry) *Service {
	return &Service{
		registry: registry,
		sources:  map[string]DataSource{},
		nextID:   1,
	}
}

func DefaultRegistry() (*Registry, error) {
	return NewRegistry(
		NewPostgresConnector(),
		NewCSVConnector(),
		NewMySQLConnector(),
		NewOracleConnector(),
		NewRESTConnector(),
	)
}

func (s *Service) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	connector, err := s.registry.Get(cfg.Type)
	if err != nil {
		return TestResult{OK: false, Message: err.Error()}, err
	}
	return connector.TestConnection(ctx, cfg)
}

func (s *Service) TestDataSource(ctx context.Context, id string) (TestResult, error) {
	source, ok := s.Get(id)
	if !ok {
		return TestResult{OK: false, Message: "data source not found"}, fmt.Errorf("data source not found")
	}

	return s.TestConnection(ctx, ConnectionConfig{
		Name:          source.Name,
		Type:          source.Type,
		ConnectionURI: source.ConnectionURI,
		Options:       source.Options,
	})
}

func (s *Service) Register(source DataSource) (DataSource, error) {
	if source.Name == "" {
		return DataSource{}, fmt.Errorf("name is required")
	}
	if source.Type == "" {
		return DataSource{}, fmt.Errorf("type is required")
	}
	if _, err := s.registry.Get(source.Type); err != nil {
		return DataSource{}, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	source.ID = fmt.Sprintf("ds_%06d", s.nextID)
	source.Status = "registered"
	s.nextID++
	s.sources[source.ID] = source

	return source, nil
}

func (s *Service) List() []DataSource {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sources := make([]DataSource, 0, len(s.sources))
	for _, source := range s.sources {
		sources = append(sources, source)
	}
	return sources
}

func (s *Service) Get(id string) (DataSource, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	source, ok := s.sources[id]
	return source, ok
}
