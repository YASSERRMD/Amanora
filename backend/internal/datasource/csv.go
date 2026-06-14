package datasource

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type CSVConnector struct{}

func NewCSVConnector() CSVConnector {
	return CSVConnector{}
}

func (CSVConnector) Type() SourceType {
	return SourceTypeCSV
}

func (CSVConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	path := cfg.ConnectionURI
	if path == "" {
		path = cfg.Options["path"]
	}
	if path == "" {
		return TestResult{OK: false, Message: "csv path is required"}, fmt.Errorf("csv path is required")
	}

	file, err := os.Open(path)
	if err != nil {
		return TestResult{OK: false, Message: err.Error()}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return TestResult{OK: false, Message: err.Error()}, err
	}

	select {
	case <-ctx.Done():
		return TestResult{OK: false, Message: ctx.Err().Error()}, ctx.Err()
	default:
		return TestResult{OK: true, Message: "csv file readable"}, nil
	}
}

func (CSVConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	path := cfg.ConnectionURI
	if path == "" {
		path = cfg.Options["path"]
	}
	if path == "" {
		return nil, fmt.Errorf("csv path is required")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	fields := make([]FieldMetadata, 0, len(header))
	for i, name := range header {
		fields = append(fields, FieldMetadata{
			Name:            name,
			OrdinalPosition: i + 1,
			DataType:        "text",
			Nullable:        true,
		})
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			if _, err := reader.Read(); err == io.EOF {
				return []AssetMetadata{{
					Name:               filepath.Base(path),
					Type:               "file",
					FullyQualifiedName: path,
					Fields:             fields,
				}}, nil
			} else if err != nil {
				return nil, err
			}
		}
	}
}
