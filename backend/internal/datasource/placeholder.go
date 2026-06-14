package datasource

import (
	"context"
	"fmt"
)

type PlaceholderConnector struct {
	sourceType SourceType
	name       string
}

func NewPlaceholderConnector(sourceType SourceType, name string) PlaceholderConnector {
	return PlaceholderConnector{sourceType: sourceType, name: name}
}

func (c PlaceholderConnector) Type() SourceType {
	return c.sourceType
}

func (c PlaceholderConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	_ = ctx
	_ = cfg
	return TestResult{OK: false, Message: c.name + " connector is not implemented yet"}, nil
}

func (c PlaceholderConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	_ = ctx
	_ = cfg
	return nil, fmt.Errorf("%s connector is not implemented yet", c.name)
}
