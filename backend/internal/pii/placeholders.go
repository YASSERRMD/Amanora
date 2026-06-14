package pii

import (
	"strings"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

type KeywordDetector struct {
	name       string
	piiType    string
	confidence float64
	keywords   []string
}

func NewKeywordDetector(name string, piiType string, confidence float64, keywords ...string) KeywordDetector {
	return KeywordDetector{name: name, piiType: piiType, confidence: confidence, keywords: keywords}
}

func (d KeywordDetector) Name() string {
	return d.name
}

func (d KeywordDetector) Detect(field datasource.FieldMetadata, samples []string) []Detection {
	_ = samples
	fieldName := strings.ToLower(field.Name)
	for _, keyword := range d.keywords {
		if strings.Contains(fieldName, strings.ToLower(keyword)) {
			return []Detection{{
				Type:       d.piiType,
				Detector:   d.name,
				Confidence: d.confidence,
				Evidence:   []string{"field name matched " + keyword},
			}}
		}
	}
	return nil
}

func NewCreditCardDetector() Detector {
	return NewKeywordDetector("credit-card-placeholder", "credit_card", 0.55, "credit_card", "card_number", "pan")
}

func DefaultDetectors() []Detector {
	return []Detector{
		NewEmailDetector(),
		NewPhoneDetector(),
		NewCreditCardDetector(),
		NewKeywordDetector("passport-placeholder", "passport", 0.55, "passport"),
		NewKeywordDetector("national-id-placeholder", "national_id", 0.55, "national_id", "ssn", "emirates_id"),
		NewKeywordDetector("name-placeholder", "name", 0.50, "full_name", "first_name", "last_name"),
		NewKeywordDetector("address-placeholder", "address", 0.50, "address", "street"),
		NewKeywordDetector("sensitive-text-placeholder", "sensitive_text", 0.45, "secret", "token", "password"),
	}
}
