package types

type RuleEngineResponse struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
