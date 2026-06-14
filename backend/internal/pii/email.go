package pii

import (
	"regexp"
	"strings"

	"github.com/YASSERRMD/Amanora/backend/internal/datasource"
)

var emailPattern = regexp.MustCompile(`(?i)\b[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}\b`)

type EmailDetector struct{}

func NewEmailDetector() EmailDetector {
	return EmailDetector{}
}

func (EmailDetector) Name() string {
	return "email-detector"
}

func (d EmailDetector) Detect(field datasource.FieldMetadata, samples []string) []Detection {
	confidence := 0.0
	evidence := []string{}

	if strings.Contains(strings.ToLower(field.Name), "email") {
		confidence = 0.72
		evidence = append(evidence, "field name contains email")
	}

	for _, sample := range samples {
		if emailPattern.MatchString(sample) {
			confidence = max(confidence, 0.98)
			evidence = append(evidence, "sample matched email pattern")
			break
		}
	}

	if confidence == 0 {
		return nil
	}

	return []Detection{{
		Type:       "email",
		Detector:   d.Name(),
		Confidence: confidence,
		Evidence:   evidence,
	}}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
