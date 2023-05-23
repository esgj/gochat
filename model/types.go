package model

import "github.com/jbrukh/bayesian"

type Intent struct {
	Class     bayesian.Class `json:"class"`
	Words     []string       `json:"words"`
	Responses []string       `json:"responses"`
	Fallback  []string       `json:"fallback"`
}
