package datasource

import (
	"context"
	"fmt"
)

type MySQLConnector struct{}

func NewMySQLConnector() MySQLConnector {
	return MySQLConnector{}
}

func (MySQLConnector) Type() SourceType {
	return SourceTypeMySQL
}

func (MySQLConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	return placeholderResult(ctx, SourceTypeMySQL)
}

func (MySQLConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, fmt.Errorf("mysql connector discovery is not implemented yet")
	}
}

func placeholderResult(ctx context.Context, sourceType SourceType) (TestResult, error) {
	select {
	case <-ctx.Done():
		return TestResult{OK: false, Message: ctx.Err().Error()}, ctx.Err()
	default:
		return TestResult{
			OK:      false,
			Message: fmt.Sprintf("%s connector is registered but not implemented yet", sourceType),
		}, nil
	}
}
