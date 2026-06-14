package policy

import "strings"

func EvaluateConditions(conditions map[string]string, facts map[string]string) bool {
	for key, expected := range conditions {
		actual, exists := facts[key]
		if expected == "$exists" {
			if !exists || strings.TrimSpace(actual) == "" {
				return false
			}
			continue
		}
		if actual != expected {
			return false
		}
	}
	return true
}
