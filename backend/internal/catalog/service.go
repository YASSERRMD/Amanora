package catalog

import (
	"fmt"
	"strings"
	"sync"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

type Asset struct {
	datasource.AssetMetadata
	ID              string   `json:"id"`
	Owner           string   `json:"owner,omitempty"`
	RiskLevel       string   `json:"riskLevel,omitempty"`
	Classifications []string `json:"classifications,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

type Field struct {
	datasource.FieldMetadata
	ID      string `json:"id"`
	AssetID string `json:"assetId"`
}

type AssetDetail struct {
	Asset
	FieldCount int `json:"fieldCount"`
}

type SearchFilter struct {
	Query          string
	Classification string
	Owner          string
	Risk           string
}

type Service struct {
	mu     sync.RWMutex
	assets map[string]Asset
	fields map[string]Field
	nextID int
}

func NewService() *Service {
	return &Service{assets: map[string]Asset{}, fields: map[string]Field{}, nextID: 1}
}

func (s *Service) UpsertAsset(asset Asset) Asset {
	s.mu.Lock()
	defer s.mu.Unlock()

	if asset.ID == "" {
		asset.ID = fmt.Sprintf("asset_%06d", s.nextID)
		s.nextID++
	}
	s.assets[asset.ID] = asset
	for index, field := range asset.Fields {
		fieldID := fmt.Sprintf("%s_field_%03d", asset.ID, index+1)
		s.fields[fieldID] = Field{FieldMetadata: field, ID: fieldID, AssetID: asset.ID}
	}
	return asset
}

func (s *Service) ListAssets(filter SearchFilter) []Asset {
	s.mu.RLock()
	defer s.mu.RUnlock()

	assets := []Asset{}
	for _, asset := range s.assets {
		if matches(asset, filter) {
			assets = append(assets, asset)
		}
	}
	return assets
}

func (s *Service) GetAsset(id string) (Asset, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	asset, ok := s.assets[id]
	return asset, ok
}

func (s *Service) GetAssetDetail(id string) (AssetDetail, bool) {
	asset, ok := s.GetAsset(id)
	if !ok {
		return AssetDetail{}, false
	}
	return AssetDetail{Asset: asset, FieldCount: len(asset.Fields)}, true
}

func (s *Service) GetField(id string) (Field, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	field, ok := s.fields[id]
	return field, ok
}

func (s *Service) AssignTags(assetID string, tags []string) (Asset, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	asset, ok := s.assets[assetID]
	if !ok {
		return Asset{}, fmt.Errorf("asset not found")
	}
	asset.Tags = unique(append(asset.Tags, tags...))
	s.assets[assetID] = asset
	return asset, nil
}

func matches(asset Asset, filter SearchFilter) bool {
	if filter.Query != "" {
		haystack := strings.ToLower(asset.Name + " " + asset.SchemaName + " " + asset.FullyQualifiedName)
		if !strings.Contains(haystack, strings.ToLower(filter.Query)) {
			return false
		}
	}
	if filter.Classification != "" && !contains(asset.Classifications, filter.Classification) {
		return false
	}
	if filter.Owner != "" && !strings.EqualFold(asset.Owner, filter.Owner) {
		return false
	}
	if filter.Risk != "" && !strings.EqualFold(asset.RiskLevel, filter.Risk) {
		return false
	}
	return true
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if strings.EqualFold(value, target) {
			return true
		}
	}
	return false
}

func unique(values []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}
