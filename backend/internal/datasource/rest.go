package datasource

import (
	"context"
	"fmt"
)

type RESTConnector struct{}

func NewRESTConnector() RESTConnector {
	return RESTConnector{}
}

func (RESTConnector) Type() SourceType {
	return SourceTypeREST
}

func (RESTConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	return placeholderResult(ctx, SourceTypeREST)
}

func (RESTConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, fmt.Errorf("rest connector discovery is not implemented yet")
	}
}
