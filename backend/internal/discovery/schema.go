package discovery

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type SchemaDiscoveryService struct{}

func NewSchemaDiscoveryService() SchemaDiscoveryService {
	return SchemaDiscoveryService{}
}

func (SchemaDiscoveryService) DiscoverSchemas(assets []datasource.AssetMetadata) []SchemaSummary {
	return SummarizeSchemas(assets)
}
