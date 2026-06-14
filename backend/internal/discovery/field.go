package discovery

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type FieldDiscoveryService struct{}

func NewFieldDiscoveryService() FieldDiscoveryService {
	return FieldDiscoveryService{}
}

func (FieldDiscoveryService) DiscoverFields(assets []datasource.AssetMetadata) []FieldSummary {
	return SummarizeFields(assets)
}
