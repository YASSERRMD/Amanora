package discovery

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type TableDiscoveryService struct{}

func NewTableDiscoveryService() TableDiscoveryService {
	return TableDiscoveryService{}
}

func (TableDiscoveryService) DiscoverTables(assets []datasource.AssetMetadata) []TableSummary {
	return SummarizeTables(assets)
}
