package kit

import (
	"encoding/json"
)

// CaseEntry delcare
type CaseEntry struct {
	Name     string
	Input    interface{}
	Expected interface{}
}

func (c CaseEntry) String() string {
	bc, _ := json.Marshal(c)
	return string(bc)
}
