package model

type Intent struct {
	Name     string   `json:"name"`
	Match    []string `json:"match"`
	Responses []string `json:"responses"`
	Fallback []string `json:"fallback"`
}

type IntentClass struct {
	Intent string   `json:"intent"`
	Words  []string `json:"words"`
}
