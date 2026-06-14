package policy

import "time"

type ExecutionResult struct {
	Decisions     []Decision `json:"decisions"`
	Compliant     bool       `json:"compliant"`
	DecisionCount int        `json:"decisionCount"`
	ExecutedAt    time.Time  `json:"executedAt"`
}

func NewExecutionResult(decisions []Decision) ExecutionResult {
	compliant := true
	for _, decision := range decisions {
		if !decision.Compliant {
			compliant = false
			break
		}
	}
	return ExecutionResult{
		Decisions:     decisions,
		Compliant:     compliant,
		DecisionCount: len(decisions),
		ExecutedAt:    time.Now().UTC(),
	}
}
