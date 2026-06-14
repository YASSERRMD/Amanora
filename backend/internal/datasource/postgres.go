package datasource

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConnector struct{}

func NewPostgresConnector() PostgresConnector {
	return PostgresConnector{}
}

func (PostgresConnector) Type() SourceType {
	return SourceTypePostgres
}

func (PostgresConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	db, err := openPostgres(cfg)
	if err != nil {
		return TestResult{OK: false, Message: err.Error()}, err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return TestResult{OK: false, Message: err.Error()}, err
	}

	return TestResult{OK: true, Message: "postgres connection ok"}, nil
}

func (PostgresConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	db, err := openPostgres(cfg)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(ctx, `
		SELECT table_schema, table_name, column_name, ordinal_position, data_type, is_nullable
		FROM information_schema.columns
		WHERE table_schema NOT IN ('pg_catalog', 'information_schema')
		ORDER BY table_schema, table_name, ordinal_position
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	assetsByName := map[string]*AssetMetadata{}
	for rows.Next() {
		var schemaName, tableName, columnName, dataType, nullable string
		var ordinal int
		if err := rows.Scan(&schemaName, &tableName, &columnName, &ordinal, &dataType, &nullable); err != nil {
			return nil, err
		}

		key := schemaName + "." + tableName
		asset, exists := assetsByName[key]
		if !exists {
			asset = &AssetMetadata{
				Name:               tableName,
				SchemaName:         schemaName,
				Type:               "table",
				FullyQualifiedName: fmt.Sprintf("%s.%s.%s", cfg.Name, schemaName, tableName),
			}
			assetsByName[key] = asset
		}

		asset.Fields = append(asset.Fields, FieldMetadata{
			Name:            columnName,
			OrdinalPosition: ordinal,
			DataType:        dataType,
			Nullable:        nullable == "YES",
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	assets := make([]AssetMetadata, 0, len(assetsByName))
	for _, asset := range assetsByName {
		assets = append(assets, *asset)
	}

	return assets, nil
}

func openPostgres(cfg ConnectionConfig) (*sql.DB, error) {
	if cfg.ConnectionURI == "" {
		return nil, fmt.Errorf("postgres connection uri is required")
	}
	return sql.Open("pgx", cfg.ConnectionURI)
}
