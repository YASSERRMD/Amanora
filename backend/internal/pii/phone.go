package pii

import (
	"regexp"
	"strings"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

var phonePattern = regexp.MustCompile(`(?i)(?:\+?\d[\d\s().-]{7,}\d)`)

type PhoneDetector struct{}

func NewPhoneDetector() PhoneDetector {
	return PhoneDetector{}
}

func (PhoneDetector) Name() string {
	return "phone-detector"
}

func (d PhoneDetector) Detect(field datasource.FieldMetadata, samples []string) []Detection {
	confidence := 0.0
	evidence := []string{}
	fieldName := strings.ToLower(field.Name)

	if strings.Contains(fieldName, "phone") || strings.Contains(fieldName, "mobile") || strings.Contains(fieldName, "tel") {
		confidence = 0.70
		evidence = append(evidence, "field name indicates phone")
	}

	for _, sample := range samples {
		if phonePattern.MatchString(sample) {
			confidence = max(confidence, 0.93)
			evidence = append(evidence, "sample matched phone pattern")
			break
		}
	}

	if confidence == 0 {
		return nil
	}

	return []Detection{{
		Type:       "phone_number",
		Detector:   d.Name(),
		Confidence: confidence,
		Evidence:   evidence,
	}}
}
