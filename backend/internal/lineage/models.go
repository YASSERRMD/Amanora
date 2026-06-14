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

func TransformationEdge(inputAssetID string, outputAssetID string, operation string) graph.Edge {
	return graph.Edge{
		FromID: inputAssetID,
		ToID:   outputAssetID,
		Type:   "transforms",
		Metadata: map[string]string{
			"model":     "transformation",
			"operation": operation,
		},
	}
}
