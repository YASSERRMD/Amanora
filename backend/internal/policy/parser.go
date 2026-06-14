package policy

import "fmt"

func ParseDocument(doc Document) (Document, error) {
	if doc.Name == "" {
		return Document{}, fmt.Errorf("policy name is required")
	}
	if doc.Category == "" {
		doc.Category = "privacy"
	}
	if doc.Severity == "" {
		doc.Severity = "medium"
	}
	if doc.Effect == "" {
		doc.Effect = "warn"
	}
	if doc.Conditions == nil {
		doc.Conditions = map[string]string{}
	}
	return doc, nil
}
