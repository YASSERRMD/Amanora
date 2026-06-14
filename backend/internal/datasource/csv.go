package datasource

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CSVConnector struct{}

func NewCSVConnector() CSVConnector {
	return CSVConnector{}
}

func (CSVConnector) Type() SourceType {
	return SourceTypeCSV
}

func (CSVConnector) TestConnection(ctx context.Context, cfg ConnectionConfig) (TestResult, error) {
	_ = ctx
	path := cfg.Options["path"]
	if path == "" {
		return TestResult{OK: false, Message: "csv path is required"}, fmt.Errorf("csv path is required")
	}

	file, err := os.Open(path)
	if err != nil {
		return TestResult{OK: false, Message: "csv file cannot be opened"}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		return TestResult{OK: false, Message: "csv header cannot be read"}, err
	}

	return TestResult{OK: true, Message: "csv file is readable"}, nil
}

func (CSVConnector) DiscoverAssets(ctx context.Context, cfg ConnectionConfig) ([]AssetMetadata, error) {
	_ = ctx
	path := cfg.Options["path"]
	if path == "" {
		return nil, fmt.Errorf("csv path is required")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	sample, err := reader.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	fields := make([]FieldMetadata, 0, len(headers))
	for index, header := range headers {
		value := ""
		if index < len(sample) {
			value = sample[index]
		}
		fields = append(fields, FieldMetadata{
			Name:            normalizeHeader(header, index),
			OrdinalPosition: index + 1,
			DataType:        inferCSVType(value),
			Nullable:        true,
		})
	}

	name := filepath.Base(path)
	return []AssetMetadata{{
		Name:               name,
		Type:               "file",
		FullyQualifiedName: fmt.Sprintf("csv.%s.%s", cfg.Name, name),
		Fields:             fields,
	}}, nil
}

func normalizeHeader(header string, index int) string {
	header = strings.TrimSpace(header)
	if header == "" {
		return fmt.Sprintf("column_%d", index+1)
	}
	return header
}

func inferCSVType(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "text"
	}
	if strings.EqualFold(value, "true") || strings.EqualFold(value, "false") {
		return "boolean"
	}
	if strings.ContainsAny(value, ".") {
		return "decimal"
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return "text"
		}
	}
	return "integer"
}
