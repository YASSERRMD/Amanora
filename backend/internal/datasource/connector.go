package datasource

import "context"

type SourceType string

const (
	SourceTypePostgres SourceType = "postgres"
	SourceTypeMySQL    SourceType = "mysql"
	SourceTypeOracle   SourceType = "oracle"
	SourceTypeCSV      SourceType = "csv"
	SourceTypeREST     SourceType = "rest"
)

type ConnectionConfig struct {
	Name          string            `json:"name"`
	Type          SourceType        `json:"type"`
	ConnectionURI string            `json:"connectionUri,omitempty"`
	Options       map[string]string `json:"options,omitempty"`
}

type AssetMetadata struct {
	Name               string          `json:"name"`
	SchemaName         string          `json:"schemaName,omitempty"`
	Type               string          `json:"type"`
	FullyQualifiedName string          `json:"fullyQualifiedName"`
	Fields             []FieldMetadata `json:"fields,omitempty"`
}

type FieldMetadata struct {
	Name            string `json:"name"`
	OrdinalPosition int    `json:"ordinalPosition"`
	DataType        string `json:"dataType"`
	Nullable        bool   `json:"nullable"`
}

type TestResult struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

type Connector interface {
	Type() SourceType
	TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error)
	DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error)
}
