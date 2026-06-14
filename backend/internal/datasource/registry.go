package datasource

import (
	"fmt"
	"sort"
	"sync"
)

type Registry struct {
	mu         sync.RWMutex
	connectors map[SourceType]Connector
}

func NewRegistry(connectors ...Connector) (*Registry, error) {
	registry := &Registry{connectors: map[SourceType]Connector{}}
	for _, connector := range connectors {
		if err := registry.Register(connector); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func (r *Registry) Register(connector Connector) error {
	if connector == nil {
		return fmt.Errorf("connector is nil")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	sourceType := connector.Type()
	if _, exists := r.connectors[sourceType]; exists {
		return fmt.Errorf("connector already registered for type %s", sourceType)
	}

	r.connectors[sourceType] = connector
	return nil
}

func (r *Registry) Get(sourceType SourceType) (Connector, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	connector, exists := r.connectors[sourceType]
	if !exists {
		return nil, fmt.Errorf("connector not registered for type %s", sourceType)
	}

	return connector, nil
}

func (r *Registry) Types() []SourceType {
	r.mu.RLock()
	defer r.mu.RUnlock()

	types := make([]SourceType, 0, len(r.connectors))
	for sourceType := range r.connectors {
		types = append(types, sourceType)
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i] < types[j]
	})

	return types
}
