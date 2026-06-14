package discovery

import "github.com/YASSERRMD/Amanora/backend/internal/datasource"

type SchemaSummary struct {
	Name       string `json:"name"`
	AssetCount int    `json:"assetCount"`
}

type TableSummary struct {
	SchemaName string `json:"schemaName"`
	Name       string `json:"name"`
	FieldCount int    `json:"fieldCount"`
}

type FieldSummary struct {
	AssetName string `json:"assetName"`
	Name      string `json:"name"`
	DataType  string `json:"dataType"`
	Nullable  bool   `json:"nullable"`
}

type SampleSet struct {
	FieldName string   `json:"fieldName"`
	Values    []string `json:"values"`
}

func SummarizeSchemas(assets []datasource.AssetMetadata) []SchemaSummary {
	counts := map[string]int{}
	for _, asset := range assets {
		schema := asset.SchemaName
		if schema == "" {
			schema = "default"
		}
		counts[schema]++
	}

	summaries := make([]SchemaSummary, 0, len(counts))
	for schema, count := range counts {
		summaries = append(summaries, SchemaSummary{Name: schema, AssetCount: count})
	}
	return summaries
}

func SummarizeTables(assets []datasource.AssetMetadata) []TableSummary {
	tables := make([]TableSummary, 0, len(assets))
	for _, asset := range assets {
		tables = append(tables, TableSummary{
			SchemaName: asset.SchemaName,
			Name:       asset.Name,
			FieldCount: len(asset.Fields),
		})
	}
	return tables
}

func SummarizeFields(assets []datasource.AssetMetadata) []FieldSummary {
	fields := []FieldSummary{}
	for _, asset := range assets {
		for _, field := range asset.Fields {
			fields = append(fields, FieldSummary{
				AssetName: asset.Name,
				Name:      field.Name,
				DataType:  field.DataType,
				Nullable:  field.Nullable,
			})
		}
	}
	return fields
}

func BuildSampleSets(assets []datasource.AssetMetadata) []SampleSet {
	samples := []SampleSet{}
	for _, asset := range assets {
		for _, field := range asset.Fields {
			samples = append(samples, SampleSet{
				FieldName: field.Name,
				Values:    []string{},
			})
		}
	}
	return samples
}
