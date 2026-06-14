package datasource

func NewOracleConnector() Connector {
	return NewPlaceholderConnector(SourceTypeOracle, "oracle")
}
