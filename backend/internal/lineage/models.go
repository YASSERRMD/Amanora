package lineage

import "github.com/YASSERRMD/Amanora/backend/internal/graph"

func IngestionEdge(sourceID string, assetID string) graph.Edge {
	return graph.Edge{
		FromID: sourceID,
		ToID:   assetID,
		Type:   "ingests",
		Metadata: map[string]string{
			"model": "ingestion",
		},
	}
}
