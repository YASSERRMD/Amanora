package datasource

import (
	"context"
	"fmt"
)

type OracleConnector struct{}

func NewOracleConnector() OracleConnector {
	return OracleConnector{}
}

func (OracleConnector) Type() SourceType {
	return SourceTypeOracle
}

func (OracleConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	return placeholderResult(ctx, SourceTypeOracle)
}

func (OracleConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, fmt.Errorf("oracle connector discovery is not implemented yet")
	}
}
